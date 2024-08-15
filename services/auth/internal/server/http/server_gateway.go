package http

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
	"github.com/rs/cors"
)

type GatewayServer struct {
	handler       http.Handler
	healthService health.HealthServiceServer
}

type ServicesServer interface {
	auth.AuthServiceServer
	health.HealthServiceServer
}

func NewGatewayServer(serviceServer ServicesServer) *GatewayServer {
	httpMux := runtime.NewServeMux()

	err := auth.RegisterAuthServiceHandlerServer(context.Background(), httpMux, serviceServer)
	if err != nil {
		log.Fatalf("Failed to register AuthService gateway: %v", err)
	}

	err = health.RegisterHealthServiceHandlerServer(context.Background(), httpMux, serviceServer)
	if err != nil {
		log.Fatalf("Failed to register HealthService gateway: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/", httpMux)
	fs := http.FileServer(http.Dir("swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081", "http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	})

	handler := corsWrapper.Handler(mux)

	return &GatewayServer{
		handler:       handler,
		healthService: serviceServer,
	}
}

func (s *GatewayServer) ListenAndServe(address string) error {
	log.Printf("Starting HTTP server on %s", address)
	return http.ListenAndServe(address, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.handler.ServeHTTP(w, r)
	}))
}
