package service

import (
	"fmt"
	"time"

	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"github.com/kurniawanxzy/backend-olshop/helper"
	"github.com/kurniawanxzy/backend-olshop/repository"
	"github.com/kurniawanxzy/backend-olshop/requests"
	"golang.org/x/crypto/bcrypt"
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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }

		data.Password = string(hashedPassword)

		if err := s.UserRepository.CreateUser(data); err != nil {
			tx.Rollback()
			return err
		}

		token, err := s.TokenVerificationRepository.GenerateToken(data.ID.String(), "email_verification")

		if err != nil {
			tx.Rollback()
			return err
		}
		helper.SendEmail(data.Email, "Email Verification", fmt.Sprintf("<h1>%s</h1>",token))

		return nil
	})
}

func (s *AuthService) VerifyUser(token, email string) error {
	if s.db == nil {
		return fmt.Errorf("missing db connection")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		tokenVerification, err:= s.TokenVerificationRepository.FindToken(token, email)

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

		if tokenVerification.User.Email != email {
			tx.Rollback()
			return fmt.Errorf("token is not for this user")
		}

		if tokenVerification.Type != "email_verification" {
			tx.Rollback()
			return fmt.Errorf("invalid token type")
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

func (s *AuthService) CreateToken(email, tokenType string) error {
	if s.db == nil {
		return fmt.Errorf("missing db connection")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		user, err := s.UserRepository.FindByEmail(email)

		if user == nil {
			tx.Rollback()
			return fmt.Errorf("user not found")
		}

		if err != nil {
			tx.Rollback()
			return err
		}

		tokenLatest, err := s.TokenVerificationRepository.FindLatestToken(user.ID.String())
		
		if err != nil {
			tx.Rollback()
			return err
		}

		if tokenLatest.ExpiredAt.After(time.Now()) {
			tx.Rollback()
			return fmt.Errorf("token still valid")
		}
		
		token, err := s.TokenVerificationRepository.GenerateToken(user.ID.String(), tokenType)

		if err != nil {
			tx.Rollback()
			return err
		}

		helper.SendEmail(user.Email, "Email Verification", fmt.Sprintf("<h1>%s</h1>",token))

		return nil
	})
}

func (s *AuthService) Login(data *requests.LoginRequest) (string, error) {
	if s.db == nil {
		return "",fmt.Errorf("missing db connection")
	}

	user, err := s.UserRepository.FindByEmail(data.Email)

	if err != nil {
		return "",err
	}

	if user == nil {
		return "",fmt.Errorf("invalid credentials")
	}

	if !user.IsVerified {
		return "",fmt.Errorf("user not verified")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return "",fmt.Errorf("invalid credentials")
	}


	token, err := helper.GenerateJWT(user)

	if err != nil {
		return "",err
	}

	return token,nil
}

func (s *AuthService) ResetPassword(data *requests.ResetPasswordRequest, user *entities.User) error {
	if s.db == nil {
		return fmt.Errorf("missing db connection")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		if user == nil {
			tx.Rollback()
			return fmt.Errorf("user not found")
		}
		
		token, err := s.TokenVerificationRepository.FindToken(data.Token, user.Email)

		if err != nil {
			tx.Rollback()
			return err
		}

		if token.IsUsed {
			tx.Rollback()
			return fmt.Errorf("token has been used")
		}

		if token.Type != "forgot_password" {
			tx.Rollback()
			return fmt.Errorf("invalid token type")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			return err
		}
		user.Password = string(hashedPassword)

		if err := s.UserRepository.UpdateUser(user); err != nil {
			tx.Rollback()
			return err
		}

		token.IsUsed = true
		if err := s.TokenVerificationRepository.UpdateToken(token); err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
}