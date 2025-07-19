package service

import (
	"app/market/internal/application/repo/ticker"
	"app/market/internal/domain"
	"app/market/internal/domain/entity"
	"app/market/internal/domain/repo"
	"app/market/internal/domain/service"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

const (
	tcpConn      = "tcp"
	generateConn = "generator"
)

type tickerConnection struct {
	soursConn domain.SourseTCPClient
	testConn  domain.TestSourseTCPClient
}

type tickerService struct {
	tickerRedisRepo repo.TickerRedisRepository
	tickerPGRepo    repo.TickerPGRepository

	tickerConnection
	mu sync.RWMutex

	currentSource     io.ReadCloser
	currentSourceName string
}

func NewTickerService(
	tikcerRepoConn repo.TickerRedisRepository,
	soursConn domain.SourseTCPClient,
	tickerPgConn repo.TickerPGRepository,
	testSoursConn domain.TestSourseTCPClient,
) service.TickerService {
	return &tickerService{
		tickerRedisRepo: tikcerRepoConn,
		tickerPGRepo:    tickerPgConn,
		tickerConnection: tickerConnection{
			soursConn: soursConn,
			testConn:  testSoursConn,
		},
		currentSource: soursConn.Sours(),
	}
}

func (t *tickerService) SwitchSourse(sourceType string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	log.Printf("Switching data source to: %s", sourceType)
	if t.currentSource != nil {
		t.currentSource.Close()
	}
	switch sourceType {
	case tcpConn:
		t.currentSource = t.soursConn.Sours()
		t.currentSourceName = tcpConn
	case generateConn:
		t.currentSource = ticker.NewGenerator()
		t.currentSourceName = generateConn
	default:
		log.Printf("Unknown source type: %s. No change.", sourceType)
	}

}

func (t *tickerService) Process(ctx context.Context, resourceMode string) error {

	var (
		buffer = make([]byte, 0, 4096) // накапливаем сюда
		temp   = make([]byte, 1024)    // читаем сюда по частям
	)

	for {

		t.mu.RLock()
		source := t.currentSource
		t.mu.RUnlock()

		n, err := source.Read(temp)
		if err != nil {

			log.Println("read error:", err)

			continue
		}

		buffer = append(buffer, temp[:n]...)

		// разбиваем по строкам
		lines := bytes.Split(buffer, []byte{'\n'})

		// обрабатываем все строки кроме последней (возможно она неполная)
		for _, line := range lines[:len(lines)-1] {
			var d entity.TickerData
			if err := json.Unmarshal(line, &d); err != nil {
				fmt.Println("parse error:", err)
				continue
			}

			err := t.tickerRedisRepo.RedisSet(ctx, &d, time.Minute)
			if err != nil {
				log.Println("redis set error:", err)
			}
		}

		buffer = lines[len(lines)-1]
	}
}

func (t *tickerService) StorePriceStats(ctx context.Context) error {
	symbols := []string{"BTCUSDT", "DOGEUSDT", "TONUSDT", "SOLUSDT", "ETHUSDT"}

	var batch []entity.AggregatedPrice

	for _, symbol := range symbols {
		avgPrice, err := t.tickerRedisRepo.GetAvgPrice(ctx, symbol)

		if err != nil {
			return err
		}

		maxPrice, err := t.tickerRedisRepo.GetMaxPrice(ctx, symbol)

		if err != nil {
			return err
		}

		minPrice, err := t.tickerRedisRepo.GetMinPrice(ctx, symbol)

		if err != nil {
			return err
		}

		data := entity.AggregatedPrice{
			PairName:     symbol,
			Exchange:     t.currentSourceName,
			Timestamp:    time.Now(),
			AveragePrice: avgPrice,
			MaxPrice:     maxPrice,
			MinPrice:     minPrice,
		}

		batch = append(batch, data)
	}

	err := t.tickerPGRepo.PutAggregatedBatch(ctx, batch)

	if err != nil {
		return err
	}

	return nil
}
