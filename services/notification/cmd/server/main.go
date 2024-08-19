package main

import (
	"log"

	"github.com/popeskul/awesome-messanger/services/notification/internal/config"
	"github.com/popeskul/awesome-messanger/services/notification/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/notification/internal/server/grpc"
	"github.com/popeskul/awesome-messanger/services/notification/internal/server/http"
	"github.com/popeskul/awesome-messanger/services/notification/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger := log.New(log.Writer(), "notification: ", log.LstdFlags)
	notificationService := services.NewNotificationService(logger)
	service := services.NewService(notificationService)

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
