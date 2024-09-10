package ports

import (
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

type ServicesServer interface {
	notification.NotificationServiceServer
	health.HealthServiceServer
}
