package service

import (
	"app/market/internal/domain"
	"app/market/internal/domain/entity"
	"app/market/internal/domain/repo"
	"app/market/internal/domain/service"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type tickerService struct {
	tickerRedisRepo repo.TickerRedisRepository
	soursConn       domain.SourseTCPClient
	tickerPGRepo    repo.TickerPGRepository
}

func NewTickerService(tikcerRepoConn repo.TickerRedisRepository, soursConn domain.SourseTCPClient, tickerPgConn repo.TickerPGRepository) service.TickerService {
	return &tickerService{
		tickerRedisRepo: tikcerRepoConn,
		soursConn:       soursConn,
		tickerPGRepo:    tickerPgConn,
	}
}

func (t *tickerService) Process(ctx context.Context, resourceMode string) error {
	conn := t.soursConn.Sours()

	var (
		buffer = make([]byte, 0, 4096) // накапливаем сюда
		temp   = make([]byte, 1024)    // читаем сюда по частям
	)

	for {

		n, err := conn.Read(temp)
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
			fmt.Println(d)
			err := t.tickerRedisRepo.RedisSet(ctx, &d, time.Minute)
			if err != nil {
				log.Println("redis set error:", err)
			}
		}

		// сохраняем последнюю строку (возможно она ещё не завершена)
		buffer = lines[len(lines)-1]
	}
}

// func (t *tickerService) StorePriceStats(ctx context.Context, symbol string) error {
// 	avgPrice, err := t.tickerRedisRepo.GetAvgPrice(ctx, symbol)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
