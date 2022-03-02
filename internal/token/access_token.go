package token

import (
	"github.com/google/uuid"
	"time"
)

//AccessToken is an interface for managing tokens
type AccessToken interface {
	// GenerateAccessToken create a new token for a specific user and durations
	GenerateAccessToken(userId uuid.UUID, duration time.Duration) (string, error)
	// GenerateRefreshToken generate a new refresh token
	GenerateRefreshToken(userId uuid.UUID , duration time.Duration)(string, error)
	// VerifyAccessToken checks if a access token is valid or not
	VerifyAccessToken(tokenString string) (bool, error)
	// VerifyRefreshToken checks if a refresh token is valid or not
	VerifyRefreshToken(tokenString string) (bool, error)
}
