package pg

import (
	"app/market/pkg/db"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

func NewBdClient(ctx context.Context, dsn string) (db.PgClient, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
