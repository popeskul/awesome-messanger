//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/popeskul/awesome-messanger/services/notification/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/notification/internal/adapters/validator"
	"github.com/popeskul/awesome-messanger/services/notification/internal/app"
	"github.com/popeskul/awesome-messanger/services/notification/internal/config"
	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/gateway_server"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/handlers"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/server"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/http/swagger"
	"github.com/popeskul/awesome-messanger/services/notification/internal/usecases"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		ProvideZapLogger,
		logger.NewZapLogger,
		usecases.NewNotificationUseCase,
		validator.ProvideProtoMessages,
		validator.NewGrpcValidator,
		handlers.NewMessageHandler,
		handlers.NewHealthHandler,
		handlers.NewHandler,
		wire.Bind(new(notification.NotificationServiceServer), new(ports.NotificationHandler)),
		wire.Bind(new(health.HealthServiceServer), new(*handlers.HealthHandler)),
		server.NewGrpcServer,
		gateway_server.NewGatewayServer,
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
