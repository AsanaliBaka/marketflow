package redise

import (
	"app/market/pkg/db"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type redisclient struct {
	masterRedise db.RedisDB
}

func NewRediseClient(ctx context.Context, redisdsn string) (db.RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisdsn,
	})

	// Проверим соединение
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	return &redisclient{
		masterRedise: &redisDb{rdb},
	}, nil
}
func (c *redisclient) RedisDB() db.RedisDB {
	return c.masterRedise
}

func (c *redisclient) Close() error {
	if c.masterRedise != nil {
		c.masterRedise.Close()
	}

	return nil
}
