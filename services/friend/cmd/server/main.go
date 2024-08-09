package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/friend/internal/server"
	"github.com/popeskul/awesome-messanger/services/friend/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger := log.New(os.Stdout, "services/friend: ", log.LstdFlags)

	validatorService := validator.New()

	service := services.NewService(logger)

	h := handlers.NewHandler(service, validatorService)

	srv := server.NewServer(cfg, h)
	log.Printf("Starting server on %s", cfg.ServerAddress)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
