package db

import (
	"context"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Handler func(ctx context.Context) error

type FullClient interface {
	PgDb() PgClient
	RedisDb() RedisClient
}

type PgClient interface {
	DB() DB
	Close() error
}
type RedisClient interface {
	RedisDB() RedisDB
	Close() error
}

type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Query struct {
	QueryRaw string
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptioin pgx.TxOptions) (pgx.Tx, error)
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest *[][]interface{}, q Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}

type RedisDB interface {
	Redis
	Pinger
	Close()
}
