package main

import (
	"app/market/internal/adapters"
	"app/market/internal/config"
	"fmt"
)

func main() {
	config, err := config.NewMultiSourceConfig()
	if err != nil {
		fmt.Println(err)
	}

	res := adapters.NewSourseConnect(config.Source1().Addr())

	res.Listen()
}
