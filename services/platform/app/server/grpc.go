package server

import (
	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
	"google.golang.org/grpc"
	"net"
)

type grpcServerWrapper struct {
	server *grpc.Server
}

func NewGRPCServer(server *grpc.Server) ports.GRPCServer {
	return &grpcServerWrapper{
		server: server,
	}
}

func (s *grpcServerWrapper) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return s.server.Serve(listener)
}

func (s *grpcServerWrapper) Stop() {
	s.server.GracefulStop()
}

func (s *grpcServerWrapper) GetGrpcServer() *grpc.Server {
	return s.server
}
