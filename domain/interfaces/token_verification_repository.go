package interfaces

import "github.com/kurniawanxzy/backend-olshop/domain/entities"


type TokenVerificationRepository interface {
	GenerateToken(userId string) string
	FindToken(token, userId string) (*entities.TokenVerification, error)
	UpdateToken(*entities.TokenVerification) error
}