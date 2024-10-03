package interfaces


type TokenVerificationRepository interface {
	GenerateToken(userId string) string
}