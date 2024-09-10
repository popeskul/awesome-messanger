package ports

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

type GRPCServer interface {
	Start(address string) error
	Stop()
	GetGrpcServer() *grpc.Server
}

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type SwaggerServer interface {
	Run() error
	Shutdown(ctx context.Context) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type ServerFactory interface {
	NewGRPCServer(registerServices func(*grpc.Server)) (GRPCServer, error)
	NewHTTPServer() (HTTPServer, error)
	NewSwaggerServer() (SwaggerServer, error)
}
