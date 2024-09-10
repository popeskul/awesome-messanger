package validator

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

func ProvideProtoMessages() []protoreflect.ProtoMessage {
	return []protoreflect.ProtoMessage{
		&notification.SendNotificationRequest{},
		&health.HealthCheckRequest{},
	}
}
