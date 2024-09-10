package ports

type UseCase interface {
	AuthUseCase() AuthUseCase
	TokenUseCase() TokenUseCase
}
