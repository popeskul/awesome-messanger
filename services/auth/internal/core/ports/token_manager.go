package ports

type TokenManager interface {
	GenerateToken(email string) (string, error)
	ValidateToken(oldToken string) (string, error)
}
