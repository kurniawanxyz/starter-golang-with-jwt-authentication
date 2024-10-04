package usecases

import (
	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"github.com/kurniawanxzy/backend-olshop/requests"
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

func (uc *AuthUseCase) CreateToken(email, tokenType string) error {
	return uc.authService.CreateToken(email, tokenType)
}

func (uc *AuthUseCase) Login(data *requests.LoginRequest) (string,error) {
	return uc.authService.Login(data)
}

func (uc *AuthUseCase) ResetPassword(data *requests.ResetPasswordRequest, user *entities.User) error {
	return uc.authService.ResetPassword(data, user)
}