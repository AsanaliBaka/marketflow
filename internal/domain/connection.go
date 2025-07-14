package domain

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PostgresClient interface {
	DB() *pgxpool.Pool
	Close() error
	Ping(ctx context.Context) error
}

type RedisClient interface {
	Ping(ctx context.Context) error
	RedisDB() *redis.Client
	Close() error
}

type SourseTCPClient interface {
	Close() error
}
