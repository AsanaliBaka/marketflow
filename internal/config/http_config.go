package config

import (
	"app/market/internal/domain"
	"app/market/pkg/cerrors"
	"net"
	"os"
)

const (
	hostEnv = "HTTP_HOST"
	portEnv = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

func NewHttpConfig() (domain.HttpConfig, error) {
	hostDataEnv := os.Getenv(hostEnv)

	if len(hostDataEnv) == 0 {
		return nil, cerrors.ErrorConfigNotFound(hostEnv)
	}

	portDataEnv := os.Getenv(portEnv)

	if len(portDataEnv) == 0 {
		return nil, cerrors.ErrorConfigNotFound(portEnv)
	}

	return &httpConfig{
		host: hostDataEnv,
		port: portDataEnv,
	}, nil

}

func (h *httpConfig) Addr() string {
	return net.JoinHostPort(h.host, h.port)
}
