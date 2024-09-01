package client_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/popeskul/awesome-messanger/services/platform/grpc/client"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func setupTestServer(t *testing.T) *grpc.Server {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()

	return s
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGrpcClient_ConnectClose(t *testing.T) {
	// Настроим тестовый сервер
	grpcServer := setupTestServer(t)
	defer grpcServer.GracefulStop()

	clientBuilder := client.NewClientBuilder().
		WithInsecure().
		WithContextDialer(bufDialer)

	grpcClient := clientBuilder.Build()

	err := grpcClient.Connect("bufnet", 5*time.Second)
	require.NoError(t, err)
	defer grpcClient.Close()

	assert.NotNil(t, grpcClient.GetConnection())

	err = grpcClient.Close()
	assert.NoError(t, err)
	assert.Nil(t, grpcClient.GetConnection())
}
