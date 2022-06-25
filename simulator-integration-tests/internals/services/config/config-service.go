package config

import (
	"os"

	"github.com/nats-io/nats.go"
)

const (
	NATS_URL = "NATS_URL"
)

type configService struct {
}

func NewConfigService() (ConfigInterface, error) {
	return &configService{}, nil
}

func (p *configService) GetNatsUrl() string {
	val := os.Getenv(NATS_URL)
	if val == "" {
		return nats.DefaultURL
	}

	return val
}
