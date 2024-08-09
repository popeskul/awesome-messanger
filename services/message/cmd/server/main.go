package main

import (
	"log"

	"github.com/popeskul/awesome-messanger/services/message/internal/config"
	"github.com/popeskul/awesome-messanger/services/message/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/message/internal/server"
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

	handler := handlers.NewHandler(service)

	srv := server.NewServer(handler)
	log.Printf("Starting server on %s", cfg.ServerAddress)

	if err := srv.ListenAndServe(cfg.ServerAddress); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
