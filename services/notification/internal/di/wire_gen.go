// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/popeskul/awesome-messanger/services/notification/internal/adapters/logger"
	"github.com/popeskul/awesome-messanger/services/notification/internal/adapters/validator"
	"github.com/popeskul/awesome-messanger/services/notification/internal/app"
	"github.com/popeskul/awesome-messanger/services/notification/internal/config"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/gateway_server"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/handlers"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/grpc/server"
	"github.com/popeskul/awesome-messanger/services/notification/internal/delivery/http/swagger"
	"github.com/popeskul/awesome-messanger/services/notification/internal/usecases"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InitializeApp() (*app.App, error) {
	configConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	zapLogger, err := ProvideZapLogger()
	if err != nil {
		return nil, err
	}
	portsLogger := logger.NewZapLogger(zapLogger)
	notificationUseCase := usecases.NewNotificationUseCase(portsLogger)
	v := validator.ProvideProtoMessages()
	portsValidator, err := validator.NewGrpcValidator(v...)
	if err != nil {
		return nil, err
	}
	notificationHandler := handlers.NewMessageHandler(notificationUseCase, portsValidator, portsLogger)
	healthHandler := handlers.NewHealthHandler()
	servicesServer := handlers.NewHandler(notificationHandler, healthHandler)
	grpcServer := server.NewGrpcServer(servicesServer)
	gatewayServer := gateway_server.NewGatewayServer(grpcServer)
	swaggerAddress := swagger.ProvideSwaggerAddress(configConfig)
	apiBaseURL := swagger.ProvideAPIBaseURL(configConfig)
	swaggerServer := swagger.NewSwaggerServer(swaggerAddress, apiBaseURL)
	appApp := app.NewApp(configConfig, portsLogger, grpcServer, gatewayServer, swaggerServer)
	return appApp, nil
}

// wire.go:

func ProvideZapLogger() (*zap.Logger, error) {
	logger2, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger2, nil
}
