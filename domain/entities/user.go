package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey;default:uuid_generate_v4()" json:"id"`
	Name       string         `gorm:"type:varchar(50);not null" json:"name" validate:"required"`
	Email      string         `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	Telp       string         `gorm:"type:varchar(15);not null" json:"telp" validate:"required,e164,min=10,max=15"`
	Password   string         `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	IsVerified bool           `gorm:"type:boolean;default:false" json:"is_verified"`
	Role       string         `gorm:"type:varchar(50);not null;default:'user'" json:"role"`
	Tokens     []TokenVerification `gorm:"foreignKey:UserID;references:ID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.New()
    return
}