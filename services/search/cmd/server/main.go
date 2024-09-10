package main

import (
	"log"

	"github.com/popeskul/awesome-messanger/services/search/internal/config"
	"github.com/popeskul/awesome-messanger/services/search/internal/graph/generated"
	"github.com/popeskul/awesome-messanger/services/search/internal/resolver"
	"github.com/popeskul/awesome-messanger/services/search/internal/server"
	"github.com/popeskul/awesome-messanger/services/search/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	userService := service.NewUserService()
	services := service.NewService(userService)

	executableSchema := generated.NewExecutableSchema(
		generated.Config{
			Resolvers: resolver.NewResolver(services),
		},
	)

	srv := server.NewServer(cfg, executableSchema)
	log.Printf("Starting server on %s", cfg.Server.HttpAddress)

	srv.Start()
}
