//go:build wireinject
// +build wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/popeskul/awesome-messanger/services/friend/internal/adapters/http/server"
	"github.com/popeskul/awesome-messanger/services/friend/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/friend/internal/app"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/friend/internal/delivery/http"
	"github.com/popeskul/awesome-messanger/services/friend/internal/swagger"
	"github.com/popeskul/awesome-messanger/services/friend/internal/usecase"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		provideZapLogger,
		wire.Bind(new(ports.Logger), new(*logger.ZapLogger)),
		logger.NewZapLogger,
		usecase.NewFriendUseCase,
		usecase.NewUseCase,
		wire.Bind(new(http.Validator), new(*validator.Validate)),
		validator.New,
		http.NewHandler,
		server.NewServer,
		swagger.NewSwaggerServer,
		app.NewApp,
	)
	return &app.App{}, nil
}

func provideZapLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
