//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/popeskul/awesome-messanger/services/message/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/message/internal/adapters/validator"
	"github.com/popeskul/awesome-messanger/services/message/internal/app"
	"github.com/popeskul/awesome-messanger/services/message/internal/config"
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/message/internal/delivery/grpc/gateway"
	"github.com/popeskul/awesome-messanger/services/message/internal/delivery/grpc/handlers"
	"github.com/popeskul/awesome-messanger/services/message/internal/delivery/grpc/server"
	"github.com/popeskul/awesome-messanger/services/message/internal/delivery/http/swagger"
	"github.com/popeskul/awesome-messanger/services/message/internal/usecases"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
	"go.uber.org/zap"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		ProvideZapLogger,
		logger.NewZapLogger,
		usecases.NewMessageUseCase,
		validator.ProvideProtoMessages,
		validator.NewGrpcValidator,
		handlers.NewMessageHandler,
		handlers.NewHealthHandler,
		handlers.NewHandler,
		wire.Bind(new(message.MessageServiceServer), new(ports.MessageHandler)),
		wire.Bind(new(health.HealthServiceServer), new(*handlers.HealthHandler)),
		server.NewGrpcServer,
		gateway.NewGatewayServer,
		swagger.ProvideSwaggerAddress,
		swagger.ProvideAPIBaseURL,
		swagger.NewSwaggerServer,
		app.NewApp,
	)
	return &app.App{}, nil
}

func ProvideZapLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
