package cron

import (
	"context"
	"sync"
	"time"
)

type Cron struct {
	mu      sync.Mutex
	tasks   []func()
	ctx     context.Context
	cancel  context.CancelFunc
	started bool
}

// New создаёт новый планировщик
func New() *Cron {
	ctx, cancel := context.WithCancel(context.Background())
	return &Cron{
		tasks:  make([]func(), 0),
		ctx:    ctx,
		cancel: cancel,
	}
}

// AddFunc добавляет задачу, которую нужно выполнять каждую минуту
func (c *Cron) AddFunc(task func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tasks = append(c.tasks, task)
}

// Start запускает выполнение задач
func (c *Cron) Start() {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()
		return
	}
	c.started = true
	c.mu.Unlock()

	go func() {
		ticker := time.NewTicker(time.Minute) // каждую минуту
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.runTasks()
			case <-c.ctx.Done():
				return
			}
		}
	}()
}

// runTasks запускает все зарегистрированные задачи в отдельных горутинах
func (c *Cron) runTasks() {
	c.mu.Lock()
	tasks := make([]func(), len(c.tasks))
	copy(tasks, c.tasks)
	c.mu.Unlock()

	for _, task := range tasks {
		go task()
	}
}

// Stop останавливает планировщик
func (c *Cron) Stop() {
	c.cancel()
	c.mu.Lock()
	c.started = false
	c.mu.Unlock()
}
