//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/popeskul/awesome-messanger/services/friend/internal/adapters/http/server"
	"github.com/popeskul/awesome-messanger/services/friend/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/friend/internal/app"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/friend/internal/delivery/http"
	"github.com/popeskul/awesome-messanger/services/friend/internal/repository"
	"github.com/popeskul/awesome-messanger/services/friend/internal/swagger"
	"github.com/popeskul/awesome-messanger/services/friend/internal/usecase"
	platformConfig "github.com/popeskul/awesome-messanger/services/platform/database/postgres/config"
	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/connection"
	platformPorts "github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

func InitializeApp(ctx context.Context) (*app.App, error) {
	wire.Build(
		config.LoadConfig,
		provideZapLogger,
		provideConnection,
		wire.Bind(new(ports.Logger), new(*logger.ZapLogger)),
		logger.NewZapLogger,
		provideFriendRepository,
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

func provideFriendRepository(conn platformPorts.Connection) ports.FriendRepository {
	repos := repository.NewRepositories(conn)
	return repos.Friend()
}

func provideConnection(ctx context.Context, cfg *config.Config) (platformPorts.Connection, error) {
	platformCfg := &platformConfig.Config{
		ConnectionString:  cfg.Database.ConnectionString,
		MaxConnections:    cfg.Database.MaxConnections,
		MinConnections:    cfg.Database.MinConnections,
		MaxConnLifetime:   cfg.Database.MaxConnLifetime,
		MaxConnIdleTime:   cfg.Database.MaxConnIdleTime,
		HealthCheckPeriod: cfg.Database.HealthCheckPeriod,
	}
	return connection.New(ctx, platformCfg)
}
