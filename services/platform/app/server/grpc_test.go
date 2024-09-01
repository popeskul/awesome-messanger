package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestNewGRPCServer(t *testing.T) {
	server := NewGRPCServer()
	assert.NotNil(t, server, "GRPCServer should not be nil")

	grpcServerWrapper, ok := server.(*grpcServerWrapper)
	assert.True(t, ok, "Server should be of type *grpcServerWrapper")
	assert.NotNil(t, grpcServerWrapper.server, "Underlying GrpcServer should not be nil")
}

func TestGRPCServerWrapper_GetGrpcServer(t *testing.T) {
	server := NewGRPCServer()
	grpcServer := server.GetGrpcServer()
	assert.NotNil(t, grpcServer, "GetGrpcServer should return a non-nil *grpc.Server")
	assert.IsType(t, &grpc.Server{}, grpcServer, "GetGrpcServer should return a *grpc.Server")
}
