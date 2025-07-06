package config

import (
	"fmt"
	"net"
	"os"
)

const (
	sours1  = "SOURS_1"
	sours2  = "SOURS_2"
	sours3  = "SOURS_3"
	hostenv = "HTTP_HOST"
)

type Sourses interface {
	ConnectSourse1() string
	ConnectSourse2() string
	ConnectSourse3() string
}
type allSourses struct {
	sourses SourseSet
}
type SourseSet struct {
	sours1 string
	sours2 string
	sours3 string
	host   string
}

func NewSourses() (Sourses, error) {
	s1 := os.Getenv(sours1)

	if len(s1) == 0 {
		return nil, fmt.Errorf("%s env not found", sours1)
	}

	s2 := os.Getenv(sours2)
	if len(s2) == 0 {
		return nil, fmt.Errorf("%s env not found", sours2)
	}
	s3 := os.Getenv(sours3)

	if len(s3) == 0 {
		return nil, fmt.Errorf("%s env not found", sours3)
	}

	host := os.Getenv(hostenv)

	if len(host) == 0 {
		return nil, fmt.Errorf("%s env not found", hostenv)

	}

	return &allSourses{
		sourses: SourseSet{
			sours1: s1,
			sours2: s2,
			sours3: s3,
			host:   host,
		},
	}, nil
}

func (s *allSourses) ConnectSourse1() string {
	return net.JoinHostPort(s.sourses.host, s.sourses.sours1)
}
func (s *allSourses) ConnectSourse2() string {
	return net.JoinHostPort(s.sourses.host, s.sourses.sours2)
}
func (s *allSourses) ConnectSourse3() string {
	return net.JoinHostPort(s.sourses.host, s.sourses.sours3)
}
