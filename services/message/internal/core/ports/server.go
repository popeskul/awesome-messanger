package ports

import (
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
)

type ServicesServer interface {
	message.MessageServiceServer
	health.HealthServiceServer
}
