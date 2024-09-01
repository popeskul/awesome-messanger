package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
	"github.com/popeskul/awesome-messanger/services/platform/config"
)

type App struct {
	config        *config.Config
	logger        ports.Logger
	serverFactory ports.ServerFactory
	grpcServer    ports.GRPCServer
	httpServer    ports.HTTPServer
	swaggerServer ports.SwaggerServer
}

func NewApp(cfg *config.Config, logger ports.Logger, serverFactory ports.ServerFactory) (*App, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}
	if serverFactory == nil {
		return nil, fmt.Errorf("serverFactory is nil")
	}

	return &App{
		config:        cfg,
		logger:        logger,
		serverFactory: serverFactory,
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	if a.config == nil {
		return fmt.Errorf("app configuration is nil")
	}

	errChan := make(chan error, 3) // Buffer for all possible server errors

	if a.config.Server.GrpcAddress != "" {
		if err := a.startGRPCServer(errChan); err != nil {
			return err
		}
	}

	if a.config.Server.GatewayAddress != "" {
		if err := a.startHTTPServer(errChan); err != nil {
			return err
		}
	}

	if a.config.Server.SwaggerAddress != "" {
		if err := a.startSwaggerServer(errChan); err != nil {
			return err
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		return err
	case <-quit:
		return a.Stop(ctx)
	case <-ctx.Done():
		return a.Stop(context.Background())
	}
}

func (a *App) startGRPCServer(errChan chan<- error) error {
	grpcServer, err := a.serverFactory.NewGRPCServer(func(s *grpc.Server) {
		// Register gRPC services here
	})
	if err != nil {
		return fmt.Errorf("failed to create gRPC server: %w", err)
	}

	go func() {
		a.logger.Info(fmt.Sprintf("Starting gRPC server on %s", a.config.Server.GrpcAddress))
		if err := grpcServer.Start(a.config.Server.GrpcAddress); err != nil {
			errChan <- fmt.Errorf("gRPC server error: %w", err)
		}
	}()

	a.grpcServer = grpcServer
	return nil
}

func (a *App) startHTTPServer(errChan chan<- error) error {
	httpServer, err := a.serverFactory.NewHTTPServer()
	if err != nil {
		return fmt.Errorf("failed to create HTTP server: %w", err)
	}

	go func() {
		a.logger.Info(fmt.Sprintf("Starting HTTP server on %s", a.config.Server.GatewayAddress))
		if err := httpServer.ListenAndServe(); err != nil {
			errChan <- fmt.Errorf("HTTP server error: %w", err)
		}
	}()

	a.httpServer = httpServer
	return nil
}

func (a *App) startSwaggerServer(errChan chan<- error) error {
	swaggerServer, err := a.serverFactory.NewSwaggerServer()
	if err != nil {
		return fmt.Errorf("failed to create Swagger server: %w", err)
	}

	go func() {
		a.logger.Info(fmt.Sprintf("Starting Swagger server on %s", a.config.Server.SwaggerAddress))
		if err := swaggerServer.Run(); err != nil {
			errChan <- fmt.Errorf("Swagger server error: %w", err)
		}
	}()

	a.swaggerServer = swaggerServer
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("Shutting down servers")

	if a.grpcServer != nil {
		a.grpcServer.Stop()
	}

	if a.httpServer != nil {
		if err := a.httpServer.Shutdown(ctx); err != nil {
			a.logger.Error(fmt.Sprintf("HTTP server shutdown error: %v", err))
		}
	}

	if a.swaggerServer != nil {
		if err := a.swaggerServer.Shutdown(ctx); err != nil {
			a.logger.Error(fmt.Sprintf("Swagger server shutdown error: %v", err))
		}
	}

	return nil
}
