package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type TokenManager struct {
	jwtSecret     []byte
	accessExpiry  time.Duration
	refreshExpiry time.Duration
	redisClient   *redis.Client
}

func NewTokenManager(secret string, accessExpiry time.Duration, refreshExpiry time.Duration, rc *redis.Client) TokenManager {
	return TokenManager {
		jwtSecret: []byte(secret),
		accessExpiry: accessExpiry,
		refreshExpiry: refreshExpiry,
		redisClient: rc,
	}
}

func (tm *TokenManager) GenerateAccessToken(uid string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"sub": uid,
		"roles": roles,
		"exp": time.Now().Add(tm.accessExpiry).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tm.jwtSecret)
}

func (tm *TokenManager) GenerateRefreshToken(uid string) (string, error) {
	// Generate crytographically secure random token
	// Store in redis with userID and expiry
	// return token
	return "", nil
}

func (tm *TokenManager) ValidateAccessToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	// Parse and validate JWT

	return nil, nil, nil
}

func (tm *TokenManager) ValidateRefreshToken(token string) (string, error) {
	// Check Redis for token and return associated uid
	return "", nil
}

func (tm *TokenManager) RevokeRefreshToken(token string) error {
	// Remove token from redis
	return nil
}

func (tm *TokenManager) RefreshTokens(refreshToken string) (string, string, error) {
	// Validate refresh token
	// Generate new access token
	// Generate new refresh token
	// Store new refresh token in Redis
	// Return both tokens
	return "", "", nil
}




