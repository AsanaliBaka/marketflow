package adapters

import (
	"app/market/internal/domain"
	"io"
	"log"
	"net"
)

type sourseTCPClient struct {
	conn      net.Conn
	soursType string
}

func NewSourseConnect(addr string) domain.SourseTCPClient {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalf("Source config error: %v", err)
	}

	return &sourseTCPClient{
		conn:      conn,
		soursType: addr,
	}

}

func (s *sourseTCPClient) Close() error {
	return s.conn.Close()
}

func (s *sourseTCPClient) Sours() io.ReadCloser {
	return s.conn
}

func (s *sourseTCPClient) SourceExchange() string {
	return s.soursType
}
