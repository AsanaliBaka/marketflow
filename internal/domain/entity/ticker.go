package entity

import "time"

type TickerData struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type AggregatedPrice struct {
	PairName     string    `db:"pair_name"`     // Название пары, например "BTCUSDT"
	Exchange     string    `db:"exchange"`      // Название биржи, например "Binance"
	Timestamp    time.Time `db:"timestamp"`     // Метка времени (конец интервала)
	AveragePrice float64   `db:"average_price"` // Средняя цена за минуту
	MinPrice     float64   `db:"min_price"`     // Минимальная цена за минуту
	MaxPrice     float64   `db:"max_price"`     // Максимальная цена за минуту
}
