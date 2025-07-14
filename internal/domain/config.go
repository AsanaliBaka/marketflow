package domain

type HttpConfig interface {
	Addr() string
}

type PgConfig interface {
	DNS() string
}

type RedisConfig interface {
	RedisConn() string
}

type SourceConfig interface {
	Host() string
	Port() string
	Addr() string
}

type MultiSourceConfig interface {
	Source1() SourceConfig
	Source2() SourceConfig
	Source3() SourceConfig
}
