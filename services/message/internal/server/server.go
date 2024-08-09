package server

import (
	"log"
	"net"

	"github.com/popeskul/awesome-messanger/services/message/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(messageServiceServer proto.MessageServiceServer) *Server {
	grpcServer := grpc.NewServer()
	proto.RegisterMessageServiceServer(grpcServer, messageServiceServer)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("message.MessageService", grpc_health_v1.HealthCheckResponse_SERVING)

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
