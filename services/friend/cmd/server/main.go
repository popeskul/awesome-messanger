package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/handlers"
	"github.com/popeskul/awesome-messanger/services/friend/internal/server"
	"github.com/popeskul/awesome-messanger/services/friend/internal/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Package handlers
// @title Friend Service API
// @version 1.0
// @description This is a Friend service API
// @host localhost:8090
// @BasePath /
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

	go func() {
		if err := srv.ListenAndServe(cfg.Server.HttpAddress); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	go func() {
		swaggerRouter := chi.NewRouter()
		swaggerRouter.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8091/swagger/doc.json"),
		))
		log.Printf("Swagger server started on %s", cfg.Server.SwaggerAddress)
		if err := http.ListenAndServe(cfg.Server.SwaggerAddress, swaggerRouter); err != nil {
			log.Fatalf("Swagger server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
