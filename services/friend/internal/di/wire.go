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
	"github.com/popeskul/awesome-messanger/services/friend/internal/service"
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
		provideRepository,
		usecase.NewFriendUseCase,
		usecase.NewUseCase,
		wire.Bind(new(http.Validator), new(*validator.Validate)),
		validator.New,
		http.NewHandler,
		server.NewServer,
		swagger.NewSwaggerServer,
		provideOutboxProcessor,
		provideService,
		app.NewApp,
	)
	return &app.App{}, nil
}

func provideZapLogger() (*zap.Logger, error) {
	return zap.NewProduction()
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

func provideRepository(conn platformPorts.Connection) ports.Repository {
	return repository.NewRepositories(conn)
}

func provideOutboxProcessor(repo ports.Repository, logger ports.Logger) ports.OutboxProcessor {
	return service.NewOutboxProcessor(repo, logger)
}

func provideService(processor ports.OutboxProcessor) ports.Service {
	return service.NewService(processor)
}
