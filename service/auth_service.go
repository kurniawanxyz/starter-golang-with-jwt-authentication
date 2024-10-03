package service

import (
	"fmt"
	"time"
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

func (s *AuthService) VerifyUser(token, userID string) error {
	if s.db == nil {
		return fmt.Errorf("missing db connection")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		tokenVerification, err:= s.TokenVerificationRepository.FindToken(token, userID)

		if err != nil {
			tx.Rollback()
			return err
		}

		if tokenVerification.ExpiredAt.Before(time.Now()) {
			tx.Rollback()
			return fmt.Errorf("token has expired")
		}

		if tokenVerification.IsUsed {
			tx.Rollback()
			return fmt.Errorf("token has been used")
		}

		if tokenVerification.UserID != userID {
			tx.Rollback()
			return fmt.Errorf("token is not for this user")
		}
		if tokenVerification.User.IsVerified {
			tx.Rollback()
			return fmt.Errorf("user already verified")
		}

		user , err := s.UserRepository.FindByEmail(tokenVerification.User.Email)
		
		if err != nil {
			tx.Rollback()
			return err
		}

		user.IsVerified = true

		if err := s.UserRepository.UpdateUser(user); err != nil {
			tx.Rollback()
			return err
		}



		tokenVerification.IsUsed = true

		if err := s.TokenVerificationRepository.UpdateToken(tokenVerification); err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
}