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
