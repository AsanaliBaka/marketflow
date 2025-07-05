package fullclient

import "app/market/pkg/db"

type fullClient struct {
	pg    db.PgClient
	redis db.RedisClient
}

func NewFullClient(pg db.PgClient, redis db.RedisClient) db.FullClient {
	return &fullClient{
		pg:    pg,
		redis: redis,
	}
}

func (f *fullClient) PgDb() db.PgClient {
	return f.pg
}

func (f *fullClient) RedisDb() db.RedisClient {
	return f.redis
}
