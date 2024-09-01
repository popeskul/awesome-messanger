// services/platform/grpc/server/middleware.go
package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			if _, ok := status.FromError(err); ok {
				return resp, err
			}
			return nil, status.Error(codes.Internal, err.Error())
		}
		return resp, nil
	}
}
