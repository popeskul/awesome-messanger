package validator

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/popeskul/awesome-messanger/services/message/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
)

func ProvideProtoMessages() []protoreflect.ProtoMessage {
	return []protoreflect.ProtoMessage{
		&message.GetMessagesRequest{},
		&message.SendMessageRequest{},
		&health.HealthCheckRequest{},
	}
}
