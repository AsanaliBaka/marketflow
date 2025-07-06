package redis

import (
	"app/market/pkg/db"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisDB struct {
	redisDb *redis.Client
}

func NewRedise(rDB *redis.Client) db.RedisDB {
	return &redisDB{
		redisDb: rDB,
	}
}

func (r *redisDB) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {

	return r.redisDb.Set(ctx, key, value, expiration).Err()
}

func (r *redisDB) Get(ctx context.Context, key string) (string, error) {
	result, err := r.redisDb.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return result, nil

}
func (r *redisDB) Del(ctx context.Context, key string) error {
	_, err := r.redisDb.Del(ctx, key).Result()

	if err != nil {
		return err
	}

	return nil

}
func (r *redisDB) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.redisDb.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return result == 1, nil
}

func (r *redisDB) Ping(ctx context.Context) error {
	return r.redisDb.Ping(ctx).Err()
}

func (r *redisDB) Close() {
	r.redisDb.Close()
}
