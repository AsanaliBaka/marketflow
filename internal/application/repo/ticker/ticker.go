package ticker

import (
	"app/market/internal/domain/entity"
	"app/market/internal/domain/repo"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type tickerRepository struct {
	redistickerRepositoryConn *redis.Client
}

func NewTickerRepository(conn *redis.Client) repo.TickerRepository {
	return &tickerRepository{
		redistickerRepositoryConn: conn,
	}
}

func (t *tickerRepository) RedisSet(ctx context.Context, data *entity.TickerData, ttl time.Duration) error {

	key := fmt.Sprintf("ticker:%s", data.Symbol)

	// Сериализуем в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal TickerData: %w", err)
	}

	// Score — это цена, Member — уникальный ID (timestamp)
	z := redis.Z{
		Score:  data.Price,
		Member: jsonData, // можно и timestamp, но json полезнее
	}

	// Добавим в Sorted Set
	err = t.redistickerRepositoryConn.ZAdd(ctx, key, z).Err()
	if err != nil {
		return fmt.Errorf("failed to ZAdd: %w", err)
	}

	// Установим TTL для ключа (если его нет)
	t.redistickerRepositoryConn.Expire(ctx, key, ttl)

	return nil
}

func (t *tickerRepository) GetMaxPrice(ctx context.Context, symbol string) (float64, error) {
	key := fmt.Sprintf("ticker:%s", symbol)

	result, err := t.redistickerRepositoryConn.ZRevRangeWithScores(ctx, key, 0, 0).Result()

	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, nil
	}

	return result[0].Score, nil
}

func (t *tickerRepository) GetMinPrice(ctx context.Context, symbol string) (float64, error) {
	key := fmt.Sprintf("ticker:%s", symbol)

	result, err := t.redistickerRepositoryConn.ZRangeWithScores(ctx, key, 0, 0).Result()

	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, nil
	}

	return result[0].Score, nil
}

func (t *tickerRepository) GetAvgPrice(ctx context.Context, symbol string) (float64, error) {
	key := fmt.Sprintf("ticker:%s", symbol)

	result, err := t.redistickerRepositoryConn.ZRangeWithScores(ctx, key, 0, -1).Result()

	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, nil
	}

	var sum float64

	for _, z := range result {
		sum += z.Score
	}

	avg := sum / float64(len(result))

	return avg, nil
}
