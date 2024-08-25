package usecases

import "github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"

type service struct {
	profileUseCase ports.ProfileUseCase
}

func NewUseCase(useCase ports.ProfileUseCase) ports.UserCase {
	return &service{
		profileUseCase: useCase,
	}
}

func (s *service) ProfileUseCase() ports.ProfileUseCase {
	return s.profileUseCase
}
