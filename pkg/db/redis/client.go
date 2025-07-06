package redis

import (
	"app/market/pkg/db"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	masterRedis db.RedisDB
}

func NewRedisClient(ctx context.Context, dsn string) (db.RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	// Проверим соединение
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	return &redisClient{
		masterRedis: &redisDB{rdb},
	}, nil
}
func (c *redisClient) RedisDB() db.RedisDB {
	return c.masterRedis
}

func (c *redisClient) Close() error {
	if c.masterRedis != nil {
		c.masterRedis.Close()
	}

	return nil
}
