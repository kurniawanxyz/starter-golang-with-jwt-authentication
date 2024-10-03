package service

import (
	"fmt"

	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"github.com/kurniawanxzy/backend-olshop/helper"
	"github.com/kurniawanxzy/backend-olshop/repository"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
	UserRepository *repository.UserRepository
	TokenVerificationRepository *repository.TokenVerificationRepository
}

func NewAuthService(db *gorm.DB, userRepository *repository.UserRepository, tokenVerificationRepository *repository.TokenVerificationRepository) *AuthService {
	return &AuthService{db, userRepository, tokenVerificationRepository}
}

func (s *AuthService) RegisterUser(data *entities.User) error {
	if s.db == nil {
		return fmt.Errorf("missing db connection")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.UserRepository.CreateUser(data); err != nil {
			tx.Rollback()
			return err
		}

		token, err := s.TokenVerificationRepository.GenerateToken(data.ID.String())

		if err != nil {
			tx.Rollback()
			return err
		}
		helper.SendEmail(data.Email, "Email Verification", fmt.Sprintf("<h1>%s</h1>",token))

		return nil
	})
}