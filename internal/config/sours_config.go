package config

import (
	"app/market/internal/domain"
	"app/market/pkg/cerrors"
	"net"
	"os"
)

const (
	soursHostEnv1 = "SOURS_HOST_1"
	soursHostEnv2 = "SOURS_HOST_2"
	soursHostEnv3 = "SOURS_HOST_3"

	soursPortEnv1 = "SOURS_PORT_1"
	soursPortEnv2 = "SOURS_PORT_2"
	soursPortEnv3 = "SOURS_PORT_3"
)

// sourceConfig - реализация SourceConfig
type sourceConfig struct {
	host string
	port string
}

func (s *sourceConfig) Host() string { return s.host }
func (s *sourceConfig) Port() string { return s.port }
func (s *sourceConfig) Addr() string { return net.JoinHostPort(s.host, s.port) }

// multiSourceConfig - реализация MultiSourceConfig
type multiSourceConfig struct {
	source1 *sourceConfig
	source2 *sourceConfig
	source3 *sourceConfig
}

func (m *multiSourceConfig) Source1() domain.SourceConfig { return m.source1 }
func (m *multiSourceConfig) Source2() domain.SourceConfig { return m.source2 }
func (m *multiSourceConfig) Source3() domain.SourceConfig { return m.source3 }

// NewMultiSourceConfig создает конфиг для всех источников
func NewMultiSourceConfig() (domain.MultiSourceConfig, error) {
	cfg1, err := parseSource(soursHostEnv1, soursPortEnv1)
	if err != nil {
		return nil, err
	}

	cfg2, err := parseSource(soursHostEnv2, soursPortEnv2)
	if err != nil {
		return nil, err
	}

	cfg3, err := parseSource(soursHostEnv3, soursPortEnv3)
	if err != nil {
		return nil, err
	}

	return &multiSourceConfig{
		source1: cfg1,
		source2: cfg2,
		source3: cfg3,
	}, nil
}

// parseSource читает host и port из env
func parseSource(hostEnv, portEnv string) (*sourceConfig, error) {
	host := os.Getenv(hostEnv)
	if len(host) == 0 {
		return nil, cerrors.ErrorConfigNotFound(hostEnv)
	}

	port := os.Getenv(portEnv)
	if len(port) == 0 {
		return nil, cerrors.ErrorConfigNotFound(portEnv)
	}

	return &sourceConfig{host: host, port: port}, nil
}
