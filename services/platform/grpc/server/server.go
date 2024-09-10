// services/platform/grpc/server/server.go
package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	grpcServer *grpc.Server
}

func NewServer(options ...grpc.ServerOption) *GrpcServer {
	server := grpc.NewServer(options...)
	reflection.Register(server)
	return &GrpcServer{grpcServer: server}
}

func (s *GrpcServer) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("Starting gRPC server on %s", address)
	return s.grpcServer.Serve(listener)
}

func (s *GrpcServer) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
}

func (s *GrpcServer) GetGrpcServer() *grpc.Server {
	return s.grpcServer
}
