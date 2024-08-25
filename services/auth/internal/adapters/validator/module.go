package validator

import (
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProvideProtoMessages() []protoreflect.ProtoMessage {
	return []protoreflect.ProtoMessage{
		&auth.LoginRequest{},
		&auth.RegisterRequest{},
		&auth.RefreshRequest{},
		&auth.LogoutRequest{},
		&health.HealthCheckRequest{},
	}
}
