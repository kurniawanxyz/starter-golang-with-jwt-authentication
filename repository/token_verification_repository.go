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

func (r *TokenVerificationRepository) GenerateToken(userId, tokenType string) (string, error) {
	var token entities.TokenVerification
	token.UserID = userId
	token.Type = tokenType
	if err := r.db.Create(&token).Error; err != nil {
		return "", err
	}
	return token.Token, nil
}

func (r *TokenVerificationRepository) FindToken(token, email string) (*entities.TokenVerification, error) {
	var tokenVerification entities.TokenVerification

    if err := r.db.Preload("User").
        Where("token = ? AND user_id = (SELECT id FROM users WHERE email = ?)", token, email).
        Order("created_at desc").
        First(&tokenVerification).Error; err != nil {
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

func (r *TokenVerificationRepository) FindLatestToken(userId string) (*entities.TokenVerification, error) {
	var tokenVerification entities.TokenVerification
	if err := r.db.Preload("User").Order("created_at desc").First(&tokenVerification, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &tokenVerification, nil
}
