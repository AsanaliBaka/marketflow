package sourse

import (
	"app/market/pkg/data"
	"bufio"
	"net"
)

type sours struct {
	conn net.Conn
}

func NewSours(conn net.Conn) data.SourseConn {
	return &sours{
		conn: conn,
	}
}

func (s *sours) Response() ([]string, error) {
	response, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil {
		return nil, err
	}

	return []string{response}, nil
}

func (s *sours) Close() {
	s.conn.Close()
}
