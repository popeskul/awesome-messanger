package middleware

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
)

// ErrorsUnaryInterceptor converts any error to an RPC error
func ErrorsUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)

		if err != nil {
			if _, ok := status.FromError(err); ok {
				return resp, err
			}

			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, domain.ErrInvalidCredentials):
				return nil, status.Error(codes.InvalidArgument, err.Error())
			default:
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
