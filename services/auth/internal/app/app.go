package app

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/popeskul/awesome-messanger/services/auth/internal/config"
	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/grpc"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/grpc_gateway"
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/http/swagger"
)

type App struct {
	Config     *config.Config
	Logger     ports.Logger
	GrpcServer *grpc.GrpcServer
	GatewaySrv *grpc_gateway.GatewayServer
	SwaggerSrv *swagger.Server
}

func NewApp(
	cfg *config.Config,
	logger ports.Logger,
	grpcServer *grpc.GrpcServer,
	gatewaySrv *grpc_gateway.GatewayServer,
	swaggerSrv *swagger.Server,
) *App {
	return &App{
		Config:     cfg,
		Logger:     logger,
		GrpcServer: grpcServer,
		GatewaySrv: gatewaySrv,
		SwaggerSrv: swaggerSrv,
	}
}

func (a *App) Run() error {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		err := a.SwaggerSrv.Run()
		if err != nil {
			a.Logger.Error("Failed to start Swagger server", "error", err)
		}

		return err
	})

	group.Go(func() error {
		if err := a.GrpcServer.Start(a.Config.Server.GrpcAddress); err != nil {
			return err
		}

		return nil
	})

	group.Go(func() error {
		if err := a.GatewaySrv.ListenAndServe(a.Config.Server.GatewayAddress); err != nil {
			return err
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("failed to start servers: %w", err)
	}

	<-ctx.Done()

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.Logger.Info("Stopping the application")

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := a.GatewaySrv.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to stop gateway server: %w", err)
		}
		return nil
	})

	group.Go(func() error {
		a.GrpcServer.Stop()
		return nil
	})

	group.Go(func() error {
		if err := a.SwaggerSrv.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to stop swagger server: %w", err)
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("error during application shutdown: %w", err)
	}

	a.Logger.Info("Application stopped")

	return nil
}
