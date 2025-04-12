package config

import "time"

type Config struct {
	// JWT Configuration
	JWTSecret          string        // Secret key for signing JWTs
	AccessTokenExpiry  time.Duration // Expiry time for access tokens (default: 15min)
	RefreshTokenExpiry time.Duration // Expiry time for refresh tokens (default: 7days)

	// Redis Configuration
	RedisURL      string // Redis connection URL
	RedisPassword string // Redis passord (optional)
	RedisDB       int    // Redis database number

	// Database Configuration
	DatabaseDriver string // Database driver (postgres, mysql, sqlite)
	DatabaseURL    string // Database connection URL

	// Email Configuration
	SMTPHost         string // SMTP server host
	SMTPPort         int    // SMTP server port
	SMTPUsername     string // SMTP username
	SMTPPassword     string // SMTP password
	EmailFrom        string // Sender email address
	EmailProvider    string // Email Provider (default: smtp)
	PasswordResetURL string // Password Reset URL
	SendGridAPIKey   string // SendGrid API Key
	MailGunDomain    string // MailGun Domain
	MailGunAPIKey    string // MailGun API Key

	// Security Settings
	PasswordMinLength int           // Minimum password length (default: 8)
	BcryptCost        int           // Bcrypt cost factor (default: 12)
	CSRFEnabled       bool          // Enable Cross-Site-Request-Forgery (default: true)
	CSRFTokenExpiry   time.Duration // Expiry time for CSRF Token (default: 1hour)
	UseHTTPS          bool          // Force HTTPS for cookies

	// Rate Limiting
	MaxLoginAttempts int           // Maximum login attempts before lockout
	LockoutDuration  time.Duration // Duration of account lockout after too many failed attempts

	// MFA Settings
	MFAEnabled bool   // Enable Multi-Factor Authentication
	MFAIssuer  string // Issuer name for TOTP

	// OAuth Settings
	GoogleClientID     string // Google OAuth client ID
	GoogleClientSecret string // Google OAuth client secret
	OAuthCallbackURL   string // Base URL for OAuth callbacks

	// Cookie settings
	CookieName     string        // Name of the auth cookie
	CookiePath     string        // Path for the auth cookie
	CookieDomain   string        // Domain for the auth cookie
	CookieMaxAge   time.Duration // Max age for the auth cookie
	CookieSecure   bool          // Use secure cookie flag
	CookieHTTPOnly bool          // Use HTTP-only cookie flag
	CookieSameSite string        // SameSite cookie policy (Lax, Strict, None)
}

func DefaultConfig() Config {
	return Config{
		JWTSecret:          "set me up",
		AccessTokenExpiry:  15 * time.Minute,
		RefreshTokenExpiry: 7 * 24 * time.Hour,
		RedisURL:           "localhost:6379",
		RedisDB:            0,
		PasswordMinLength:  8,
		BcryptCost:         12,
		UseHTTPS:           true,
		MaxLoginAttempts:   5,
		LockoutDuration:    15 * time.Minute,
		EmailProvider:      "smtp",
		SMTPHost:           "localhost",
		SMTPPort:           25,
		CSRFEnabled:        true,
		CSRFTokenExpiry:    1 * time.Hour,
		MFAEnabled:         false,
		MFAIssuer:          "Gauth",
		CookieName:         "gauth_token",
		CookiePath:         "/",
		CookieMaxAge:       15 * time.Minute,
		CookieSecure:       true,
		CookieHTTPOnly:     true,
		CookieSameSite:     "lax",
	}
}
