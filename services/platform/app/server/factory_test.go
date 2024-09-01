package server_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
	"github.com/popeskul/awesome-messanger/services/platform/app/server"
	"github.com/popeskul/awesome-messanger/services/platform/config"
)

func TestNewServerFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockSwaggerHandler := func(w http.ResponseWriter, r *http.Request) {}
	cfg := &config.Config{
		HTTPAddress:    ":8080",
		SwaggerAddress: ":8081",
	}

	factory := server.NewServerFactory(cfg, mockLogger, mockSwaggerHandler)

	assert.NotNil(t, factory, "ServerFactory should not be nil")
	assert.IsType(t, &server.ServerFactory{}, factory, "ServerFactory should be of correct type")
}

func TestServerFactory_NewGRPCServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockSwaggerHandler := func(w http.ResponseWriter, r *http.Request) {}
	cfg := &config.Config{}

	factory := server.NewServerFactory(cfg, mockLogger, mockSwaggerHandler)

	grpcServer, err := factory.NewGRPCServer()

	assert.NoError(t, err, "NewGRPCServer should not return an error")
	assert.NotNil(t, grpcServer, "GRPC Server should not be nil")
}

func TestServerFactory_NewHTTPServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockSwaggerHandler := func(w http.ResponseWriter, r *http.Request) {}
	cfg := &config.Config{
		HTTPAddress: ":8080",
	}

	factory := server.NewServerFactory(cfg, mockLogger, mockSwaggerHandler)

	// Set up logger expectations
	mockLogger.EXPECT().Info(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes()

	httpServer, err := factory.NewHTTPServer()

	assert.NoError(t, err, "NewHTTPServer should not return an error")
	assert.NotNil(t, httpServer, "HTTP Server should not be nil")
}

func TestServerFactory_NewSwaggerServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockSwaggerHandler := func(w http.ResponseWriter, r *http.Request) {}
	cfg := &config.Config{
		SwaggerAddress: ":8081",
	}

	factory := server.NewServerFactory(cfg, mockLogger, mockSwaggerHandler)

	swaggerServer, err := factory.NewSwaggerServer()

	assert.NoError(t, err, "NewSwaggerServer should not return an error")
	assert.NotNil(t, swaggerServer, "Swagger Server should not be nil")
}
