package usecases

import (
	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"github.com/kurniawanxzy/backend-olshop/service"
)

type AuthUseCase struct {
	authService *service.AuthService
}

func NewAuthUseCase(authService *service.AuthService) *AuthUseCase {
	return &AuthUseCase{authService}
}

func (uc *AuthUseCase) RegisterUser(data *entities.User) error {
	return uc.authService.RegisterUser(data)
}

func (uc *AuthUseCase) VerifyUser(token, userID string) error {
	return uc.authService.VerifyUser(token, userID)
}

func (uc *AuthUseCase) CreateToken(email string) error {
	return uc.authService.CreateToken(email)
}