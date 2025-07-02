package config

import (
	"fmt"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type PgConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PgConfig, error) {
	dsn := os.Getenv(dsnEnvName)

	if len(dsn) == 0 {
		return nil, fmt.Errorf("%s env not set", dsnEnvName)
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (p *pgConfig) DSN() string {
	return p.dsn
}
