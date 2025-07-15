package main

import (
	"app/market/internal/adapters"
	"app/market/internal/application/repo/ticker"
	"app/market/internal/application/service"
	"app/market/internal/config"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	conf, err := config.NewMultiSourceConfig()
	if err != nil {
		fmt.Println(err)
	}

	redisConf, err := config.NewRedisConfig()
	if err != nil {
		fmt.Println(err)
	}

	app := adapters.NewSourseConnect(conf.Source1().Addr())

	redisConn := adapters.NewRedisClient(redisConf.RedisConn(), ctx)
	tickerRepo := ticker.NewTickerRepository(redisConn.RedisDB())
	aplication := service.NewTickerService(tickerRepo, app)

	go func() {
		if err := aplication.Process(ctx, ""); err != nil {
			fmt.Println("process error:", err)
		}
	}()

	// Блокировка main, чтобы программа не завершилась
	select {}

}
