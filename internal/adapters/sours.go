package adapters

import (
	"app/market/internal/domain"
	"log"
	"net"
)

type sourseTCPClient struct {
	conn net.Conn
}

func NewSourseConnect(addr string) domain.SourseTCPClient {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalf("Source config error: %v", err)
	}

	return &sourseTCPClient{
		conn: conn,
	}

}

func (s *sourseTCPClient) Close() error {
	return s.conn.Close()
}

func (s *sourseTCPClient) Sours() net.Conn {
	return s.conn
}
