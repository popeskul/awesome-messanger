package ports

import (
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
)

type MessageHandler interface {
	message.MessageServiceServer
}

type Handlers interface {
	MessageHandler() MessageHandler
}
