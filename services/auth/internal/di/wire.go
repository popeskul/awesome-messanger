//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/popeskul/awesome-messanger/services/auth/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/auth/internal/adapters/token"
	"github.com/popeskul/awesome-messanger/services/auth/internal/adapters/validator"
	"github.com/popeskul/awesome-messanger/services/auth/internal/app"
	"github.com/popeskul/awesome-messanger/services/auth/internal/config"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/grpc"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/grpc_gateway"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/handlers"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/http/swagger"
	"github.com/popeskul/awesome-messanger/services/auth/internal/usecases"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		ProvideZapLogger,
		logger.NewZapLogger,
		token.NewTokenService,
		usecases.NewAuthUseCase,
		usecases.NewTokenUseCase,
		validator.ProvideProtoMessages,
		validator.NewGrpcValidator,
		handlers.NewAuthHandler,
		handlers.NewHealthHandler,
		wire.Bind(new(health.HealthServiceServer), new(*handlers.HealthHandler)),
		handlers.NewHandler,
		grpc.NewServer,
		grpc_gateway.NewGatewayServer,
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
