package ports

import (
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

type ProfileHandler interface {
	profile.ProfileServiceServer
}

type Handlers interface {
	ProfileHandler
}
