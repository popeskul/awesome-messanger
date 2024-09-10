package ports

type Token interface {
	GenerateToken(email string) (string, error)
	ValidateToken(oldToken string) (string, error)
}
