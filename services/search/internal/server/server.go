package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/popeskul/awesome-messanger/services/search/internal/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config, executableSchema graphql.ExecutableSchema) *Server {
	srv := handler.NewDefaultServer(executableSchema)

	http.Handle("/graphql", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() {
	log.Printf("connect to http://localhost%s/ for GraphQL playground", s.cfg.Server.HttpAddress)
	if err := http.ListenAndServe(s.cfg.Server.HttpAddress, nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
