package validator

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

func ProvideProtoMessages() []protoreflect.ProtoMessage {
	return []protoreflect.ProtoMessage{
		&profile.CreateProfileRequest{},
		&profile.UpdateProfileRequest{},
		&profile.GetProfileRequest{},
		&health.HealthCheckRequest{},
	}
}
