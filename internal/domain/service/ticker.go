package service

import "context"

type TickerService interface {
	Process(ctx context.Context, resourceMode string) error
}

type DataSource interface {
	Read(p []byte) (n int, err error)
	Close() error
}
