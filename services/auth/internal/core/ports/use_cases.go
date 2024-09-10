package ports

type UseCases interface {
	AuthUseCase() AuthUseCase
}
