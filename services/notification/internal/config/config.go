package config

import (
	"errors"
	"os"
)

type Config struct {
	ServerAddress string
}

func LoadConfig() (*Config, error) {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		return nil, errors.New("SERVER_ADDRESS is required")
	}

	return &Config{
		ServerAddress: serverAddress,
	}, nil
}
