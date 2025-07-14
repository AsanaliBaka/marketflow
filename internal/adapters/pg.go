package adapters

import (
	"app/market/internal/domain"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresClient struct {
	dbc *pgxpool.Pool
}

func NewPgClient(dsn string, ctx context.Context) domain.PostgresClient {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	}

	return &postgresClient{
		dbc: dbc,
	}
}

func (p *postgresClient) DB() *pgxpool.Pool {
	return p.dbc
}

func (p *postgresClient) Close() error {
	if p.dbc != nil {
		p.dbc.Close()
	}

	return nil
}

func (p *postgresClient) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}
