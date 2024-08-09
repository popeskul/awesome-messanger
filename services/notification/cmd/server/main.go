package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/popeskul/awesome-messanger/services/notification/internal/config"
	"github.com/popeskul/awesome-messanger/services/notification/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/notification/internal/server"
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
	validationService := validator.New()

	handler := handlers.NewHandler(service, validationService)

	srv := server.NewServer(handler)
	log.Printf("Starting server on %s", cfg.ServerAddress)

	if err := srv.ListenAndServe(cfg.ServerAddress); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
