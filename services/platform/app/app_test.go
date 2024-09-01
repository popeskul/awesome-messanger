package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
	"github.com/popeskul/awesome-messanger/services/platform/config"
)

func TestNewApp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockServerFactory := ports.NewMockServerFactory(ctrl)
	cfg := &config.Config{}

	app, err := NewApp(cfg, mockLogger, mockServerFactory)

	assert.NoError(t, err)
	assert.NotNil(t, app)
}

func TestApp_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockServerFactory := ports.NewMockServerFactory(ctrl)
	mockGRPCServer := ports.NewMockGRPCServer(ctrl)
	mockHTTPServer := ports.NewMockHTTPServer(ctrl)
	mockSwaggerServer := ports.NewMockSwaggerServer(ctrl)

	cfg := &config.Config{
		GRPCAddress:    ":50051",
		HTTPAddress:    ":8080",
		SwaggerAddress: ":8081",
	}

	mockServerFactory.EXPECT().NewGRPCServer().Return(mockGRPCServer, nil)
	mockServerFactory.EXPECT().NewHTTPServer().Return(mockHTTPServer, nil)
	mockServerFactory.EXPECT().NewSwaggerServer().Return(mockSwaggerServer, nil)

	mockLogger.EXPECT().Info(gomock.Any()).AnyTimes()

	// Use channels to control when each server's Start or ListenAndServe method returns
	grpcStartCh := make(chan struct{})
	httpStartCh := make(chan struct{})
	swaggerStartCh := make(chan struct{})

	mockGRPCServer.EXPECT().Start(cfg.GRPCAddress).DoAndReturn(func(string) error {
		<-grpcStartCh
		return nil
	})
	mockHTTPServer.EXPECT().ListenAndServe().DoAndReturn(func() error {
		<-httpStartCh
		return nil
	})
	mockSwaggerServer.EXPECT().ListenAndServe().DoAndReturn(func() error {
		<-swaggerStartCh
		return nil
	})

	// Expect Stop and Shutdown calls
	mockGRPCServer.EXPECT().Stop()
	mockHTTPServer.EXPECT().Shutdown(gomock.Any()).Return(nil)
	mockSwaggerServer.EXPECT().Shutdown(gomock.Any()).Return(nil)

	app, _ := NewApp(cfg, mockLogger, mockServerFactory)

	// Create a context with cancel
	ctx, cancel := context.WithCancel(context.Background())

	// Start the app in a goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Start(ctx)
	}()

	// Give some time for the servers to "start"
	time.Sleep(100 * time.Millisecond)

	// Close the start channels to simulate servers starting successfully
	close(grpcStartCh)
	close(httpStartCh)
	close(swaggerStartCh)

	// Cancel the context to stop the app
	cancel()

	// Wait for the app to stop
	err := <-errCh
	assert.NoError(t, err)
}

func TestApp_Stop(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockServerFactory := ports.NewMockServerFactory(ctrl)
	mockGRPCServer := ports.NewMockGRPCServer(ctrl)
	mockHTTPServer := ports.NewMockHTTPServer(ctrl)
	mockSwaggerServer := ports.NewMockSwaggerServer(ctrl)

	cfg := &config.Config{
		GRPCAddress:    ":50051",
		HTTPAddress:    ":8080",
		SwaggerAddress: ":8081",
	}

	mockLogger.EXPECT().Info(gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Error(gomock.Any()).AnyTimes()

	mockGRPCServer.EXPECT().Stop()
	mockHTTPServer.EXPECT().Shutdown(gomock.Any()).Return(nil)
	mockSwaggerServer.EXPECT().Shutdown(gomock.Any()).Return(nil)

	app, _ := NewApp(cfg, mockLogger, mockServerFactory)
	appInstance := app.(*App)
	appInstance.grpcServer = mockGRPCServer
	appInstance.httpServer = mockHTTPServer
	appInstance.swaggerServer = mockSwaggerServer

	err := app.Stop(context.Background())
	assert.NoError(t, err)
}

func TestApp_Start_Errors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockServerFactory := ports.NewMockServerFactory(ctrl)

	testCases := []struct {
		name          string
		config        *config.Config
		setupMocks    func()
		expectedError string
	}{
		{
			name: "GRPC server creation fails",
			config: &config.Config{
				GRPCAddress: ":50051",
			},
			setupMocks: func() {
				mockServerFactory.EXPECT().NewGRPCServer().Return(nil, errors.New("GRPC server creation failed"))
			},
			expectedError: "failed to create gRPC server: GRPC server creation failed",
		},
		{
			name: "HTTP server creation fails",
			config: &config.Config{
				HTTPAddress: ":8080",
			},
			setupMocks: func() {
				mockServerFactory.EXPECT().NewHTTPServer().Return(nil, errors.New("HTTP server creation failed"))
			},
			expectedError: "failed to create HTTP server: HTTP server creation failed",
		},
		{
			name: "Swagger server creation fails",
			config: &config.Config{
				SwaggerAddress: ":8081",
			},
			setupMocks: func() {
				mockServerFactory.EXPECT().NewSwaggerServer().Return(nil, errors.New("Swagger server creation failed"))
			},
			expectedError: "failed to create Swagger server: Swagger server creation failed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()
			app, _ := NewApp(tc.config, mockLogger, mockServerFactory)
			err := app.Start(context.Background())
			assert.EqualError(t, err, tc.expectedError)
		})
	}
}
