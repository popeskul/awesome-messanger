package grpc

import (
	"log"
	"net"

	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerGRPC struct {
	grpcServer *grpc.Server
}

type ServicesServer interface {
	notification.NotificationServiceServer
	health.HealthServiceServer
}

func NewGrpcServer(servicesServer ServicesServer) *ServerGRPC {
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.MaxSendMsgSize(10*1024*1024))
	notification.RegisterNotificationServiceServer(grpcServer, servicesServer)
	health.RegisterHealthServiceServer(grpcServer, servicesServer)

	reflection.Register(grpcServer)

	return &ServerGRPC{
		grpcServer: grpcServer,
	}
}

func (s *ServerGRPC) ListenAndServe(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on %s", address)
	return s.grpcServer.Serve(listener)
}
