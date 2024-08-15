package main

import (
	"log"

	"github.com/popeskul/awesome-messanger/services/auth/internal/config"
	"github.com/popeskul/awesome-messanger/services/auth/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/auth/internal/server/grpc"
	"github.com/popeskul/awesome-messanger/services/auth/internal/server/http"
	"github.com/popeskul/awesome-messanger/services/auth/internal/services"
	"github.com/popeskul/awesome-messanger/services/auth/internal/swagger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger := log.New(log.Writer(), "auth: ", log.LstdFlags)

	authService := services.NewAuthService(logger)
	tokenService := services.NewTokenService(logger)
	service := services.NewService(tokenService, authService)

	handler, err := handlers.NewHandlers(service)
	if err != nil {
		log.Fatalf("Error creating handlers: %v", err)
	}

	grpcServer := grpc.NewGrpcServer(handler)
	gatewayServer := http.NewGatewayServer(handler)

	go func() {
		swaggerServer := swagger.NewSwaggerServer(cfg.Server.SwaggerAddress, "http://"+cfg.Server.HttpAddress)
		if err := swaggerServer.Run(); err != nil {
			log.Fatalf("Failed to start Swagger UI server: %v", err)
		}
	}()

	go func() {
		if err := grpcServer.ListenAndServe(cfg.Server.GrpcAddress); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	if err := gatewayServer.ListenAndServe(cfg.Server.HttpAddress); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
