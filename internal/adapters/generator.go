package adapters

import (
	"app/market/internal/domain"
	"app/market/internal/domain/service"
	"io"
)

type generatorClinet struct {
	reader service.DataSource
}

func NewGeneratorClient(reader service.DataSource) domain.TestSourseTCPClient {
	return &generatorClinet{
		reader: reader,
	}
}

func (g *generatorClinet) Close() error {
	return g.reader.Close()
}

func (g *generatorClinet) Sours() io.ReadCloser {
	return g.reader
}

func (g *generatorClinet) SourceExchange() string {
	return "test exchange"
}
