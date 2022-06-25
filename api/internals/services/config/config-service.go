package config

import (
	"os"
	"strconv"

	"github.com/nats-io/nats.go"
)

const (
	PORT     = "PORT"
	NATS_URL = "NATS_URL"
)

// configService implements ConfigInterface contract
type configService struct {
}

// NewConfigService create a concerete instance of ConfigInterface interface
func NewConfigService() (ConfigInterface, error) {
	return &configService{}, nil
}

// GetListeningPort returns the port for api server to listen to
func (p *configService) GetListeningPort() int {
	val := os.Getenv(PORT)
	if val == "" {
		return 8080
	}

	port, err := strconv.Atoi(val)
	if err != nil {
		return 8080
	}

	return port
}

// GetNatsUrl returns the uri that can be used to connect to NATS
func (p *configService) GetNatsUrl() string {
	val := os.Getenv(NATS_URL)
	if val == "" {
		return nats.DefaultURL
	}

	return val
}
