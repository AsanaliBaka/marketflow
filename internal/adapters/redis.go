package adapters

import (
	"app/market/internal/domain"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type redisConn struct {
	client *redis.Client
}

func NewRedisClient(redisDSN string, ctx context.Context) domain.RedisClient {
	opt, _ := redis.ParseURL(redisDSN)

	rdb := redis.NewClient(opt)

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	return &redisConn{
		client: rdb,
	}

}

func (r *redisConn) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisConn) RedisDB() *redis.Client {
	return r.client
}

func (r *redisConn) Close() error {
	return r.client.Close()
}
