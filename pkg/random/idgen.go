package random

import (
	"app/market/internal/domain/entity"
	"math/rand"
	"time"
)

func RandomData() *entity.TickerData {
	symbols := []string{"BTCUSDT", "DOGEUSDT", "TONUSDT", "SOLUSDT", "ETHUSDT"}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomSymbol := symbols[r.Intn(len(symbols))]
	randomPrice := 100 + r.Float64()*(50000-100)

	return &entity.TickerData{
		Symbol:    randomSymbol,
		Price:     randomPrice,
		Timestamp: time.Now().Unix(),
	}
}
