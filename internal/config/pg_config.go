package config

import (
	"app/market/internal/domain"
	"app/market/pkg/cerrors"
	"os"
)

const (
	pgEnv = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewPgDSN() (domain.PgConfig, error) {
	pgEnvData := os.Getenv(pgEnv)

	if len(pgEnvData) == 0 {
		return nil, cerrors.ErrorConfigNotFound(pgEnv)
	}

	return &pgConfig{
		dsn: pgEnvData,
	}, nil

}

func (p *pgConfig) DNS() string {
	return p.dsn
}
