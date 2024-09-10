package server_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/popeskul/awesome-messanger/services/platform/grpc/server"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func setupServer(t *testing.T) *server.GrpcServer {
	lis = bufconn.Listen(bufSize)
	s := server.NewServer(grpc.UnaryInterceptor(server.UnaryErrorInterceptor()))
	go func() {
		if err := s.GetGrpcServer().Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	return s
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGrpcServer_StartStop(t *testing.T) {
	server := setupServer(t)
	defer server.Stop()

	require.NotNil(t, server)

	conn, err := grpc.Dial("bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	assert.NotNil(t, conn)
}
