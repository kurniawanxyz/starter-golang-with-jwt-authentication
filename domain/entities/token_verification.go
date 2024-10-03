package entities

import (
	"crypto/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenVerification struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    string    `gorm:"type:char(36);not null" json:"user_id"`
	Token     string    `gorm:"type:varchar(255);not null" json:"token"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	ExpiredAt time.Time `gorm:"type:timestamp;not null" json:"expired_at"`
	IsUsed    bool      `gorm:"type:boolean;default:false" json:"is_used"`
	Type 	  string 	`gorm:"type:enum('email_verification', 'forgot_password');not null;default:'email_verification'" json:"type"`
	
	User      User 	`gorm:"foreignKey:UserID;references:ID"`
}

func (token *TokenVerification) BeforeCreate(tx *gorm.DB) (err error) {
	token.ID = uuid.New()
	token.ExpiredAt = time.Now().Add(time.Minute * 15) // Set expired token in 24 hours
	token.Token, err = generateRandomString(6) // Generate a random token with 32 bytes
	return
}


func generateRandomString(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		bytes[i] = charset[bytes[i]%byte(len(charset))]
	}
	return string(bytes), nil
}