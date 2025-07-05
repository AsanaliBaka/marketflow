package config

import (
	"fmt"
	"os"
)

const (
	sours1 = "SOURS_1"
	sours2 = "SOURS_2"
	sours3 = "SOURS_3"
)

type Sourses interface {
	ConnectSourse() SourseSet
}
type allSourses struct {
	sourses SourseSet
}
type SourseSet struct {
	sours1 string
	sours2 string
	sours3 string
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

	return &allSourses{
		sourses: SourseSet{
			sours1: s1,
			sours2: s2,
			sours3: s3,
		},
	}, nil
}

func (s *allSourses) ConnectSourse() SourseSet {
	return s.sourses
}
