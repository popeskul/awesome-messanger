package main

import (
	"log"

	"github.com/popeskul/awesome-messanger/services/message/internal/config"
	"github.com/popeskul/awesome-messanger/services/message/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/message/internal/server/grpc"
	"github.com/popeskul/awesome-messanger/services/message/internal/server/http"
	"github.com/popeskul/awesome-messanger/services/message/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger := log.New(log.Writer(), "message: ", log.LstdFlags)
	messageService := services.NewMessageService(logger)
	service := services.NewService(messageService)

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
