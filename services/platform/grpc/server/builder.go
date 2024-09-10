// services/platform/grpc/server/builder.go
package server

import (
	"google.golang.org/grpc"
)

type Builder struct {
	options []grpc.ServerOption
}

func NewServerBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithUnaryInterceptor(interceptor grpc.UnaryServerInterceptor) *Builder {
	b.options = append(b.options, grpc.UnaryInterceptor(interceptor))
	return b
}

func (b *Builder) WithOptions(options ...grpc.ServerOption) *Builder {
	b.options = append(b.options, options...)
	return b
}

func (b *Builder) Build() *GrpcServer {
	return NewServer(b.options...)
}
