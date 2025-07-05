package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1. Подключение с таймаутом
	conn, err := net.DialTimeout("tcp", "localhost:40102", time.Millisecond)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer conn.Close()

	// 5. Выводим и завершаем работу

}
