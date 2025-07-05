package app

import (
	"app/market/internal/config"
	"app/market/pkg/closer"
	"app/market/pkg/db"
	"app/market/pkg/db/fullclient"
	"app/market/pkg/db/pg"
	"app/market/pkg/db/redise"
	"context"
	"log"
)

type serviceProvider struct {
	pgConfig      config.PgConfig
	httpConfig    config.HTTPConfig
	soursesConfig config.Sourses
	redisConfig   config.RedisConfig

	dbClient  db.FullClient
	txManager db.TxManager
}

func (s *serviceProvider) PGConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SoursConfig() config.Sourses {
	if s.soursesConfig == nil {
		cfg, err := config.NewSourses()
		if err != nil {
			log.Fatalf("failed to get sours config: %s", err.Error())
		}

		s.soursesConfig = cfg
	}

	return s.soursesConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.FullClient {
	if s.dbClient == nil {
		pgClient, err := pg.NewBdClient(ctx, s.PGConfig().DSN())

		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = pgClient.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(pgClient.Close)

		redisClient, err := redise.NewRediseClient(ctx, s.redisConfig.DSN())

		if err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}

		err = redisClient.RedisDB().Ping(ctx)

		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(redisClient.Close)

		fullcl := fullclient.NewFullClient(pgClient, redisClient)

		return fullcl

	}

	return s.dbClient
}
