package redise

import (
	"app/market/pkg/db"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisDb struct {
	redisDb *redis.Client
}

func NewRedise(rDB *redis.Client) db.RedisDB {
	return &redisDb{
		redisDb: rDB,
	}
}

func (r *redisDb) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {

	return r.redisDb.Set(ctx, key, value, expiration).Err()
}

func (r *redisDb) Get(ctx context.Context, key string) (string, error) {
	result, err := r.redisDb.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return result, nil

}
func (r *redisDb) Del(ctx context.Context, key string) error {
	_, err := r.redisDb.Del(ctx, key).Result()

	if err != nil {
		return err
	}

	return nil

}
func (r *redisDb) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.redisDb.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return result == 1, nil
}

func (r *redisDb) Ping(ctx context.Context) error {
	return r.redisDb.Ping(ctx).Err()
}

func (r *redisDb) Close() {
	r.redisDb.Close()
}
