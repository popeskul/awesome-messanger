package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"

	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	_ "github.com/popeskul/awesome-messanger/services/friend/swagger"
)

type Server struct {
	httpServer *http.Server
	handler    http.Handler
	logger     ports.Logger
}

func NewServer(cfg *config.Config, handler ports.HandlerFriends, logger ports.Logger) *Server {
	r := chi.NewRouter()

	r.Use(LoggingMiddleware(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8011", "http://localhost:8010"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler

	r.Use(corsHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/add-friend", handler.PostAddFriend)
		r.Get("/friends", handler.GetFriends)
		r.Post("/respond-friend-request", handler.PostRespondFriendRequest)
		r.Get("/live", handler.GetLive)
		r.Get("/ready", handler.GetReady)
	})

	// Serve Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", cfg.Server.GatewayAddress)),
	))

	// Serve Swagger JSON
	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		swaggerJSONPath := filepath.Join(".", "swagger", "swagger.json")
		logger.Info(fmt.Sprintf("Serving Swagger JSON from: %s", swaggerJSONPath))
		http.ServeFile(w, r, swaggerJSONPath)
	})

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Server.GatewayAddress,
			Handler: r,
		},
		handler: r,
		logger:  logger,
	}
}

func (s *Server) ListenAndServe() error {
	s.logger.Info(fmt.Sprintf("Starting HTTP server on %s", s.httpServer.Addr))
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server error: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		s.logger.Info("Shutting down HTTP server")
		return s.httpServer.Shutdown(ctx)
	})
	if err := group.Wait(); err != nil {
		return fmt.Errorf("failed to shutdown servers: %w", err)
	}

	<-ctx.Done()
	return nil
}
