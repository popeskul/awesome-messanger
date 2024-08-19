package http

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

type GatewayServer struct {
	httpMux *runtime.ServeMux
}

type ServicesServer interface {
	notification.NotificationServiceServer
	health.HealthServiceServer
}

func NewGatewayServer(serviceServer ServicesServer) *GatewayServer {
	httpMux := runtime.NewServeMux()

	err := notification.RegisterNotificationServiceHandlerServer(context.Background(), httpMux, serviceServer)
	if err != nil {
		log.Fatalf("Failed to register MessageService gateway: %v", err)
	}

	err = health.RegisterHealthServiceHandlerServer(context.Background(), httpMux, serviceServer)
	if err != nil {
		log.Fatalf("Failed to register HealthService gateway: %v", err)
	}

	return &GatewayServer{
		httpMux: httpMux,
	}
}

func (s *GatewayServer) ListenAndServe(address string) error {
	log.Printf("Starting HTTP server on %s", address)
	return http.ListenAndServe(address, s.httpMux)
}
