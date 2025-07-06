package sourse

import (
	"app/market/pkg/data"
	"context"
	"net"
)

type soursClient struct {
	s data.SourseConn
}

func NewSourseClient(ctx context.Context, sourse string) (data.SourseClient, error) {
	conn, err := net.Dial("tcp", sourse)

	if err != nil {
		return nil, err
	}

	return &soursClient{
		s: &sours{conn: conn},
	}, nil

}

func (s *soursClient) SourseConn() data.SourseConn {
	return s.s
}

func (s *soursClient) Close() error {
	if s.s != nil {
		s.s.Close()
	}

	return nil
}
