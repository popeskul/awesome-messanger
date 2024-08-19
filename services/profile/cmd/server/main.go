package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/popeskul/awesome-messanger/services/profile/internal/config"
	"github.com/popeskul/awesome-messanger/services/profile/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/profile/internal/server"
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
	validatorService := validator.New()

	handler := handlers.NewHandler(service, validatorService)

	srv := server.NewServer(handler)
	log.Printf("Starting server on %s", cfg.ServerAddress)

	if err := srv.ListenAndServe(cfg.ServerAddress); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
