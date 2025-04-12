package models

import (
	"time"

	"gorm.io/gorm"
)

type VerificationCode struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string    `gorm:"index;not null"`
	Code      string    `gorm:"not null"`
	Type      string    `gorm:"not null"` // "email", "password_reset"
	ExpiresAt time.Time `gorm:"not null"`
	UsedAt    *time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
