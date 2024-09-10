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
	GrpcAddress    string
	SwaggerAddress string
}

func LoadConfig() (*Config, error) {
	GatewayAddress := os.Getenv("SERVER_GATEWAY_ADDRESS")
	GrpcAddress := os.Getenv("SERVER_GRPC_ADDRESS")
	SwaggerAddress := os.Getenv("SERVER_SWAGGER_ADDRESS")

	if GatewayAddress == "" {
		return nil, errors.New("SERVER_GATEWAY_ADDRESS is required")
	}

	if GrpcAddress == "" {
		return nil, errors.New("SERVER_GRPC_ADDRESS is required")
	}

	if SwaggerAddress == "" {
		return nil, errors.New("SERVER_SWAGGER_ADDRESS is required")
	}

	return &Config{
		Server: Server{
			GatewayAddress: GatewayAddress,
			GrpcAddress:    GrpcAddress,
			SwaggerAddress: SwaggerAddress,
		},
	}, nil
}
