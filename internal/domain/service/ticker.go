package service

import "context"

type TickerService interface {
	Process(ctx context.Context, resourceMode string) error
}
