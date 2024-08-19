package server

import (
	"log"
	"net"

	"github.com/popeskul/awesome-messanger/services/profile/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(profileServiceServer proto.ProfileServiceServer) *Server {
	grpcServer := grpc.NewServer()
	proto.RegisterProfileServiceServer(grpcServer, profileServiceServer)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("profile.ProfileService", grpc_health_v1.HealthCheckResponse_SERVING)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) ListenAndServe(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on %s", address)
	return s.grpcServer.Serve(listener)
}
