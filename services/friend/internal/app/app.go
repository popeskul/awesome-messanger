package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"

	"github.com/popeskul/awesome-messanger/services/friend/internal/adapters/http/server"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/friend/internal/swagger"
)

type App struct {
	cfg           *config.Config
	httpServer    *server.Server
	swaggerServer *swagger.Server
	service       ports.Service
	Logger        ports.Logger
}

func NewApp(
	cfg *config.Config,
	logger ports.Logger,
	httpServer *server.Server,
	swaggerServer *swagger.Server,
	service ports.Service,
) *App {
	return &App{
		cfg:           cfg,
		httpServer:    httpServer,
		swaggerServer: swaggerServer,
		Logger:        logger,
		service:       service,
	}
}

func (a *App) Run() error {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return a.httpServer.ListenAndServe()
	})

	group.Go(func() error {
		return a.swaggerServer.Run()
	})

	group.Go(func() error {
		a.service.OutboxProcessor().Start(ctx)
		return nil
	})

	if err := group.Wait(); err != nil {
		a.Logger.Error("App failed", "error", err)
		return fmt.Errorf("app failed: %w", err)
	}

	<-ctx.Done()
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.Logger.Info("Stopping the application")

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return a.httpServer.Shutdown(ctx)
	})

	group.Go(func() error {
		return a.swaggerServer.Shutdown(ctx)
	})

	group.Go(func() error {
		a.service.OutboxProcessor().Stop()
		return nil
	})

	return group.Wait()
}
