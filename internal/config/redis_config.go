package config

import (
	"fmt"
	"os"
)

const (
	dsnEnvRedis = "REDIS_DSN"
)

type RedisConfig interface {
	DSN() string
}

type redisConfig struct {
	dsn string
}

func NewRedisConfig() (RedisConfig, error) {
	dsn := os.Getenv(dsnEnvRedis)

	if len(dsn) == 0 {
		return nil, fmt.Errorf("%s env not set", dsnEnvRedis)
	}

	return &redisConfig{
		dsn: dsn,
	}, nil
}

func (r *redisConfig) DSN() string {
	return r.dsn
}
