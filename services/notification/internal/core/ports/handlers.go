package ports

import (
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

type NotificationHandler interface {
	notification.NotificationServiceServer
}

type Handlers interface {
	NotificationHandler
}
