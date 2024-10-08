package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/middleware"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

type GrpcServer struct {
	server          *grpc.Server
	ServicesServer  ports.ServicesServer
	registeredFuncs []func(*grpc.Server)
}

func NewGrpcServer(servicesServer ports.ServicesServer) *GrpcServer {
	return &GrpcServer{
		ServicesServer: servicesServer,
		registeredFuncs: []func(*grpc.Server){
			func(s *grpc.Server) { notification.RegisterNotificationServiceServer(s, servicesServer) },
			func(s *grpc.Server) { health.RegisterHealthServiceServer(s, servicesServer) },
		},
	}
}

func (s *GrpcServer) RegisterService(f func(*grpc.Server)) {
	s.registeredFuncs = append(s.registeredFuncs, f)
}

func (s *GrpcServer) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ErrorsUnaryInterceptor()),
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.MaxSendMsgSize(10*1024*1024),
	)

	for _, f := range s.registeredFuncs {
		f(s.server)
	}

	reflection.Register(s.server)

	log.Printf("Starting gRPC server on %s", address)
	return s.server.Serve(listener)
}

func (s *GrpcServer) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}
