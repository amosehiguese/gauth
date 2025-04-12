package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID           string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID       string         `gorm:"index;not null"`
	RefreshToken string         `gorm:"uniqueIndex;not null"`
	UserAgent    string
	IP           string
	ExpiresAt    time.Time      `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}