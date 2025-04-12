package models

import (
	"time"
)

type AuditLog struct {
	ID        string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    *string `gorm:"index"`
	EventType string  `gorm:"not null"`
	IP        string
	UserAgent string
	Metadata  string
	CreatedAt time.Time
}

// constants for audit log event types
const (
	AuditEventLogin             = "login"
	AuditEventLoginFailed       = "login_failed"
	AuditEventLogout            = "logout"
	AuditEventRegistration      = "registration"
	AuditEventEmailVerification = "email_verification"
	AuditEventPasswordChange    = "password_change"
	AuditEventPasswordReset     = "password_reset"
	AuditEventRoleAssigned      = "role_assigned"
	AuditEventRoleRevoked       = "role_revoked"
	AuditEventMFAEnabled        = "mfa_enabled"
	AuditEventMFADisabled       = "mfa_disabled"
	AuditEventAccountLocked     = "account_locked"
	AuditEventAccountUnlocked   = "account_unlocked"
	AuditEventSessionRevoked    = "session_revoked"
	AuditEventUserDeleted       = "user_deleted"
)
