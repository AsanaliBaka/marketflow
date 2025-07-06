package cron

import "time"

type Scheduler interface {
	Start()
	Stop()
}
type scheduler struct {
	task   func()
	ticker *time.Ticker
	done   chan struct{}
}

func NewScheduler(interval time.Duration, task func()) Scheduler {
	return &scheduler{
		task:   task,
		ticker: time.NewTicker(interval),
		done:   make(chan struct{}),
	}
}

func (s *scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				s.task()
			case <-s.done:
				return
			}
		}
	}()
}

func (s *scheduler) Stop() {
	s.ticker.Stop()
	close(s.done)
}
