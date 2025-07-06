package app

import (
	"app/market/internal/config"
	"app/market/pkg/closer"
	"app/market/pkg/data"
	"app/market/pkg/data/sourse"
	"app/market/pkg/db"
	"app/market/pkg/db/fullclient"
	"app/market/pkg/db/pg"
	"app/market/pkg/db/redis"
	"app/market/pkg/db/transaction"
	"context"
	"log"
)

type serviceProvider struct {
	pgConfig      config.PgConfig
	httpConfig    config.HTTPConfig
	soursesConfig config.Sourses
	redisConfig   config.RedisConfig

	soursClient []data.SourseClient
	dbClient    db.FullClient
	txManager   db.TxManager
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
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
		pgClient, err := pg.NewPostgresClient(ctx, s.PGConfig().DSN())

		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = pgClient.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(pgClient.Close)

		redisClient, err := redis.NewRedisClient(ctx, s.redisConfig.DSN())

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

func (s *serviceProvider) SoursClient(ctx context.Context) []data.SourseClient {
	if s.soursClient == nil {
		s1, err := sourse.NewSourseClient(ctx, s.soursesConfig.ConnectSourse1())

		if err != nil {
			log.Fatalf("failed create connection %s", s.soursesConfig.ConnectSourse1())
		}

		closer.Add(s1.Close)

		s2, err := sourse.NewSourseClient(ctx, s.soursesConfig.ConnectSourse2())

		if err != nil {
			log.Fatalf("failed create connection %s", s.soursesConfig.ConnectSourse2())
		}

		closer.Add(s2.Close)

		s3, err := sourse.NewSourseClient(ctx, s.soursesConfig.ConnectSourse3())

		if err != nil {
			log.Fatalf("failed create connection %s", s.soursesConfig.ConnectSourse3())
		}

		closer.Add(s3.Close)

		return []data.SourseClient{s1, s2, s3}

	}

	return s.soursClient
}
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).PgDb().DB())
	}

	return s.txManager
}
