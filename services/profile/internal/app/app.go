package app

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/popeskul/awesome-messanger/services/profile/internal/config"
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/grpc/gateway_server"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/grpc/server"
	"github.com/popeskul/awesome-messanger/services/profile/internal/delivery/http/swagger"
)

type App struct {
	Config     *config.Config
	Logger     ports.Logger
	GrpcServer *server.GrpcServer
	GatewaySrv *gateway_server.GatewayServer
	SwaggerSrv *swagger.Server
}

func NewApp(
	cfg *config.Config,
	logger ports.Logger,
	grpcServer *server.GrpcServer,
	gatewaySrv *gateway_server.GatewayServer,
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
		return fmt.Errorf("failed to start app: %w", err)
	}

	<-ctx.Done()

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.Logger.Info("Stopping app")
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.Logger.Info("Stopping gRPC server")
		a.GrpcServer.Stop()
		return nil
	})

	group.Go(func() error {
		a.Logger.Info("Stopping gateway server")
		return a.GatewaySrv.Shutdown(ctx)
	})

	group.Go(func() error {
		a.Logger.Info("Stopping Swagger server")
		return a.SwaggerSrv.Shutdown(ctx)
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("failed to stop app: %w", err)
	}

	return nil
}
