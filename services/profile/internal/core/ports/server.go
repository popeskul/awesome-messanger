package ports

import (
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

type ServicesServer interface {
	profile.ProfileServiceServer
	health.HealthServiceServer
}
