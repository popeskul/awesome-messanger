package config

import (
	"errors"
	"os"
)

type Config struct {
	Server Server
}

type Server struct {
	GatewayAddress string
	SwaggerAddress string
}

func LoadConfig() (*Config, error) {
	HttpAddress := os.Getenv("SERVER_GATEWAY_ADDRESS")
	SwaggerAddress := os.Getenv("SERVER_SWAGGER_ADDRESS")

	if HttpAddress == "" {
		return nil, errors.New("SERVER_HTTP_ADDRESS is required")
	}

	if SwaggerAddress == "" {
		return nil, errors.New("SERVER_SWAGGER_ADDRESS is required")
	}

	return &Config{
		Server: Server{
			GatewayAddress: HttpAddress,
			SwaggerAddress: SwaggerAddress,
		},
	}, nil
}
