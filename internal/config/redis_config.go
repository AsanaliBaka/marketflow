package config

import (
	"app/market/internal/domain"
	"app/market/pkg/cerrors"
	"os"
)

const (
	redisEnv = "REDIS_DSN"
)

type redisConfig struct {
	redisConn string
}

func NewRedisConfig() (domain.RedisConfig, error) {
	redisDNS := os.Getenv(redisEnv)

	if len(redisDNS) == 0 {
		return nil, cerrors.ErrorConfigNotFound(redisEnv)
	}

	return &redisConfig{
		redisConn: redisDNS,
	}, nil
}

func (r *redisConfig) RedisConn() string {
	return r.redisConn
}
