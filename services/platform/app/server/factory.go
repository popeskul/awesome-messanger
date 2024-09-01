package server

import (
	"net/http"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
	"github.com/popeskul/awesome-messanger/services/platform/app/swagger"
	"github.com/popeskul/awesome-messanger/services/platform/config"
	platformGrpc "github.com/popeskul/awesome-messanger/services/platform/grpc/server"
	"google.golang.org/grpc"
)

type ServerFactory struct {
	config         *config.Config
	logger         ports.Logger
	swaggerHandler http.HandlerFunc
}

func NewServerFactory(cfg *config.Config, logger ports.Logger, swaggerHandler http.HandlerFunc) ports.ServerFactory {
	return &ServerFactory{
		config:         cfg,
		logger:         logger,
		swaggerHandler: swaggerHandler,
	}
}

func (f *ServerFactory) NewGRPCServer(registerServices func(*grpc.Server)) (ports.GRPCServer, error) {
	interceptor := platformGrpc.UnaryErrorInterceptor()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor),
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.MaxSendMsgSize(10*1024*1024),
	)

	registerServices(grpcServer)

	return NewGRPCServer(grpcServer), nil
}

func (f *ServerFactory) NewHTTPServer() (ports.HTTPServer, error) {
	return NewHTTPServer(f.config.Server.GatewayAddress, f.logger)
}

func (f *ServerFactory) NewSwaggerServer() (ports.SwaggerServer, error) {
	return swagger.NewSwaggerServer(f.config.Server.SwaggerAddress, f.swaggerHandler), nil
}
