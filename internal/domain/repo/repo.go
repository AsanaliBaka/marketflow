package repo

import (
	"app/market/internal/domain/entity"
	"context"
	"time"
)

type TickerRedisRepository interface {
	RedisSet(ctx context.Context, data *entity.TickerData, ttl time.Duration) error
	GetMaxPrice(ctx context.Context, symbol string) (float64, error)
	GetMinPrice(ctx context.Context, symbol string) (float64, error)
	GetAvgPrice(ctx context.Context, symbol string) (float64, error)
}

type TickerPGRepository interface {
	PutAggregatedBatch(ctx context.Context, batch []entity.AggregatedPrice) error
}
