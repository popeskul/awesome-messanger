package ports

type TokenUseCase interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) (string, error)
}
