package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                 string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email              string         `gorm:"uniqueIndex;not null" json:"email"`
	EmailVerified      bool           `gorm:"default:false" json:"email_verified"`
	PasswordHash       string         `gorm:"not null" json:"-"`
	FirstName          *string        `gorm:"size:100" json:"first_name"`
	LastName           *string        `gorm:"size:100" json:"last_name"`
	PhoneNumber        *string        `gorm:"size:20" json:"phone_number"`
	PhotoURL           *string        `json:"photo_url"`
	MFAEnabled         bool           `gorm:"default:false" json:"mfa_enabled"`
	MFASecret          string         `gorm:"type:text" json:"-"`
	FailedLoginCount   int            `gorm:"default:0" json:"-"`
	LastFailedLogin    *time.Time     `json:"-"`
	AccountLocked      bool           `gorm:"default:false" json:"account_locked"`
	AccountLockedUntil *time.Time     `json:"-"`
	LastLoginAt        *time.Time     `json:"last_login_at"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Metadata           *string        `gorm:"type:text" json:"metadata"`

	Roles       []Role       `gorm:"many2many:user_roles;" json:"roles"`
	Permissions []Permission `gorm:"many2many:user_permissions;" json:"permissions"`
}

// SafeUser returns a User object with sensitive fields removed
func (u *User) SafeUser() map[string]interface{} {
	var firstName, lastName, phoneNumber, photoURL string
	if u.FirstName != nil {
		firstName = *u.FirstName
	}

	if u.LastName != nil {
		lastName = *u.LastName
	}

	if u.PhoneNumber != nil {
		phoneNumber = *u.PhotoURL
	}

	var lastLoginAt time.Time
	if u.LastLoginAt != nil {
		lastLoginAt = *u.LastLoginAt
	}

	return map[string]interface{}{
		"id":             u.ID,
		"email":          u.Email,
		"email_verified": u.EmailVerified,
		"first_name":     firstName,
		"last_name":      lastName,
		"phone_number":   phoneNumber,
		"photo_url":      photoURL,
		"mfa_enabled":    u.MFAEnabled,
		"account_locked": u.AccountLocked,
		"last_login_at":  lastLoginAt,
		"created_at":     u.CreatedAt,
		"roles":          u.Roles,
		"permissions":    u.Permissions,
	}
}

// HasRole checks if the user has specific role
func (u *User) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

// HasPermission checks if the user has a specific permission
func (u *User) HasPermission(permissionName string) bool {
	for _, permission := range u.Permissions {
		if permission.Name == permissionName {
			return true
		}
	}

	for _, role := range u.Roles {
		for _, permission := range role.Permissions {
			if permission.Name == permissionName {
				return true
			}
		}
	}

	return false
}

// isLocked returns whether the account is currently locked
func (u *User) IsLocked() bool {
	if !u.AccountLocked {
		return false
	}

	if u.AccountLockedUntil == nil {
		return true // No unlock time means permanently locked
	}

	return time.Now().Before(*u.AccountLockedUntil)
}
