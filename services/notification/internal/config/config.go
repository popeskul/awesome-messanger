package config

import (
	"errors"
	"os"
)

type Config struct {
	Server Server
}

type Server struct {
	HttpAddress string
	GrpcAddress string
}

func LoadConfig() (*Config, error) {
	HttpAddress := os.Getenv("SERVER_HTTP_ADDRESS")
	GrpcAddress := os.Getenv("SERVER_GRPC_ADDRESS")

	if HttpAddress == "" {
		return nil, errors.New("SERVER_HTTP_ADDRESS is required")
	}

	if GrpcAddress == "" {
		return nil, errors.New("SERVER_GRPC_ADDRESS is required")
	}

	return &Config{
		Server: Server{
			HttpAddress: HttpAddress,
			GrpcAddress: GrpcAddress,
		},
	}, nil
}
