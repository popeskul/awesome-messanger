package config

import (
	"errors"
	"os"
)

// Config holds the configuration values for the application
type Config struct {
	ServerAddress string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		return nil, errors.New("SERVER_ADDRESS is required")
	}

	return &Config{
		ServerAddress: serverAddress,
	}, nil
}
