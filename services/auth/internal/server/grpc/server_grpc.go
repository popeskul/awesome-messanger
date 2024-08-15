package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerGRPC struct {
	grpcServer *grpc.Server
}

type ServicesServer interface {
	auth.AuthServiceServer
	health.HealthServiceServer
}

func NewGrpcServer(servicesServer ServicesServer) *ServerGRPC {
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.MaxSendMsgSize(10*1024*1024))
	auth.RegisterAuthServiceServer(grpcServer, servicesServer)
	health.RegisterHealthServiceServer(grpcServer, servicesServer)

	reflection.Register(grpcServer)

	return &ServerGRPC{
		grpcServer: grpcServer,
	}
}

func (s *ServerGRPC) ListenAndServe(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on %s", address)
	return s.grpcServer.Serve(listener)
}
