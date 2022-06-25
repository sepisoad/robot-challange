package config

// ConfigInterface defines the contracts for a configuration service
type ConfigInterface interface {
	GetListeningPort() int
	GetNatsUrl() string
}
