package swagger

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	_ "github.com/popeskul/awesome-messanger/services/friend/swagger"
)

type Server struct {
	address    string
	apiBaseURL string
	server     *http.Server
	logger     ports.Logger
}

func NewSwaggerServer(cfg *config.Config, logger ports.Logger) *Server {
	return &Server{
		address:    cfg.Server.SwaggerAddress,
		apiBaseURL: fmt.Sprintf("http://%s", cfg.Server.GatewayAddress),
		logger:     logger,
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8011", "http://0.0.0.0:8011"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsMiddleware.Handler)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", s.address)),
	))

	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		swaggerJSONPath := filepath.Join(".", "swagger", "swagger.json")
		s.logger.Info(fmt.Sprintf("Serving Swagger JSON from: %s", swaggerJSONPath))
		http.ServeFile(w, r, swaggerJSONPath)
	})

	s.server = &http.Server{
		Addr:    s.address,
		Handler: r,
	}

	s.logger.Info("Swagger server is running on", "address", s.address)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		s.logger.Info("Shutting down Swagger server")
		return s.server.Shutdown(ctx)
	}
	return nil
}
