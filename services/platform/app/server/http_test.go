package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
)

func TestNewHTTPServer(t *testing.T) {
	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock logger
	mockLogger := ports.NewMockLogger(ctrl)

	// Set up expectations for the logger
	// Expect Info to be called with 7 arguments (1 string + 6 any)
	mockLogger.EXPECT().Info(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes()

	// Create the HTTP server
	server, err := NewHTTPServer(":8080", mockLogger)
	assert.NoError(t, err, "NewHTTPServer should not return an error")
	assert.NotNil(t, server, "HTTP server should not be nil")

	// Test health endpoint
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	// Use the Handler field of the http.Server
	httpServer, ok := server.(*http.Server)
	assert.True(t, ok, "server should be of type *http.Server")
	httpServer.Handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Health check should return 200 OK")
	assert.Equal(t, "OK", rr.Body.String(), "Health check should return 'OK'")
}
