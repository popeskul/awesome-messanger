package swagger

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSwaggerServer(t *testing.T) {
	tests := []struct {
		name           string
		customHandler  SwaggerHandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Default Handler",
			customHandler:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "Swagger UI",
		},
		{
			name: "Custom Handler",
			customHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Custom Swagger UI"))
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Custom Swagger UI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := NewSwaggerServer(":8082", tt.customHandler)
			assert.NoError(t, err)
			assert.NotNil(t, server)

			req, err := http.NewRequest("GET", "/swagger/", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			server.(*swaggerServer).server.Handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.expectedBody)
		})
	}
}

func TestSwaggerServer_ListenAndServe(t *testing.T) {
	customHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test Swagger"))
	}

	// Find an available port
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()

	server, err := NewSwaggerServer(fmt.Sprintf(":%d", port), customHandler)
	require.NoError(t, err)

	// Start the server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			t.Errorf("Unexpected error from ListenAndServe: %v", err)
		}
	}()

	// Wait for the server to start
	var resp *http.Response
	for i := 0; i < 10; i++ {
		resp, err = http.Get(fmt.Sprintf("http://localhost:%d/swagger/", port))
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()

	// Shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	assert.NoError(t, err)
}
