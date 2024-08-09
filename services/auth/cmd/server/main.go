package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/popeskul/awesome-messanger/services/auth/internal/config"
	"github.com/popeskul/awesome-messanger/services/auth/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/auth/internal/server"
	"github.com/popeskul/awesome-messanger/services/auth/internal/services"
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
	validation := validator.New()

	authServiceServer := handlers.NewHandlers(service, validation)
	srv := server.NewServer(authServiceServer)
	log.Printf("Starting gRPC server on %s", cfg.ServerAddress)

	if err := srv.ListenAndServe(cfg.ServerAddress); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
