package main

import (
	"log"

	"github.com/popeskul/awesome-messanger/services/auth/internal/config"
	"github.com/popeskul/awesome-messanger/services/auth/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	srv := server.NewServer(cfg)
	log.Printf("Starting server on %s", cfg.ServerAddress)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
