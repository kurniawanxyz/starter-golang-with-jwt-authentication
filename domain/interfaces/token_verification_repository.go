package interfaces

import "github.com/kurniawanxzy/backend-olshop/domain/entities"


type TokenVerificationRepository interface {
	GenerateToken(userId string, tokeType string) string
	FindToken(token, email string) (*entities.TokenVerification, error)
	UpdateToken(*entities.TokenVerification) error
	FindLatestToken(userId string) (*entities.TokenVerification, error)
}