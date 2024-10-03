package repository

import (
	"github.com/kurniawanxzy/backend-olshop/domain/entities"
	"gorm.io/gorm"
)

type TokenVerificationRepository struct {
	db *gorm.DB
}

func NewTokenVerificationRepository(db *gorm.DB) *TokenVerificationRepository {
	return &TokenVerificationRepository{db}
}

func (r *TokenVerificationRepository) GenerateToken(userId string) (string, error) {
	var token entities.TokenVerification
	token.UserID = userId
	if err := r.db.Create(&token).Error; err != nil {
		return "", err
	}
	return token.Token, nil
}

func (r *TokenVerificationRepository) FindToken(token, userId string) (*entities.TokenVerification, error) {
	var tokenVerification entities.TokenVerification
	if err := r.db.Preload("User").First(&tokenVerification,"token = ?", token).Error; err != nil {
		return nil, err
	}
	return &tokenVerification, nil
}

func (r *TokenVerificationRepository) UpdateToken(token *entities.TokenVerification) error {
	if err := r.db.Save(&token).Error; err != nil {
		return err
	}
	return nil
}