package main

import (
	"log"
	"os"

	"github.com/popeskul/awesome-messanger/services/profile/internal/config"
	"github.com/popeskul/awesome-messanger/services/profile/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/profile/internal/server/grpc"
	"github.com/popeskul/awesome-messanger/services/profile/internal/server/http"
	"github.com/popeskul/awesome-messanger/services/profile/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger := log.New(os.Stdout, "profile", log.LstdFlags)
	profileService := services.NewProfileService(logger)
	service := services.NewService(profileService)

	handler, err := handlers.NewHandler(service)
	if err != nil {
		log.Fatalf("Error creating handlers: %v", err)
	}

	grpcServer := grpc.NewGrpcServer(handler)
	gatewayServer := http.NewGatewayServer(handler)

	go func() {
		if err := grpcServer.ListenAndServe(cfg.Server.GrpcAddress); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	if err := gatewayServer.ListenAndServe(cfg.Server.HttpAddress); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
