//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/popeskul/awesome-messanger/services/profile/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/profile/internal/adapters/validator"
	"github.com/popeskul/awesome-messanger/services/profile/internal/app"
	"github.com/popeskul/awesome-messanger/services/profile/internal/config"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/grpc/gateway_server"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/grpc/handlers"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/grpc/server"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/http/swagger"
	"github.com/popeskul/awesome-messanger/services/profile/internal/usecases"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/health"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		ProvideZapLogger,
		logger.NewZapLogger,
		usecases.NewProfileUseCase,
		validator.ProvideProtoMessages,
		validator.NewGrpcValidator,
		handlers.NewMessageHandler,
		handlers.NewHealthHandler,
		wire.Bind(new(health.HealthServiceServer), new(*handlers.HealthHandler)),
		handlers.NewHandler,
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
