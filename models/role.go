package models

import (
	"time"

	"gorm.io/gorm"
)

// Role represents a role in the RBAC system
type Role struct {
	ID          string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"update_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Permissions []Permission   `gorm:"many2many:role_permissions;" json:"permissions"`
	Users       []User         `gorm:"many2many:user_roles;" json:"-"`
}

// Permission represents a permission in the RBAC system
type Permission struct {
	ID          string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"update_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Roles       []Role         `gorm:"many2many:role_permissions;" json:"-"`
	Users       []User         `gorm:"many2many:user_permissions;" json:"-"`
}
