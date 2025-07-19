package ticker

import (
	"app/market/internal/domain/service"
	"app/market/pkg/random"
	"encoding/json"
	"fmt"
	"time"
)

type generator struct {
}

func NewGenerator() service.DataSource {
	return &generator{}
}

func (g *generator) Read(p []byte) (n int, err error) {

	data := random.RandomData()

	if data == nil {
		return 0, fmt.Errorf("failed to generate random data: %w", err)

	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	jsonData = append(jsonData, '\n') // Добавляем разделитель
	copy(p, jsonData)

	time.Sleep(1 * time.Second)
	return len(jsonData), nil

}

func (g *generator) Close() error {

	return nil
}
