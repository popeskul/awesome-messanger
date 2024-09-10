package ports

import (
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
)

type MessageHandler interface {
	message.MessageServiceServer
	//GetMessages(ctx context.Context, req *message.GetMessagesRequest) (*message.GetMessagesResponse, error)
	//SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error)
}

type Handlers interface {
	MessageHandler() MessageHandler
}
