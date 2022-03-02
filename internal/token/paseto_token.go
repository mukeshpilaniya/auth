package token

import (
	"errors"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

// PasetoToken is a paseto token manager
type PasetoToken struct {
	paseto    *paseto.V2
	secretKey []byte
}

// GenerateAccessToken create a new token for a specific username and duration
func (p *PasetoToken) GenerateAccessToken(userId uuid.UUID, duration time.Duration) (string, error) {
	token, err := NewToken(userId, duration)

	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.secretKey, token, nil)
}

// VerifyAccessToken checks if the token is valid or not
func (p *PasetoToken) VerifyAccessToken(tokenString string) (bool, error) {
	token := &Token{}

	err := p.paseto.Decrypt(tokenString, p.secretKey, token, nil)

	if err != nil {
		return false, ErrInvalidToken
	}
	err = token.Valid()

	if err != nil {
		return false, err
	}
	return true, nil
}

// GenerateRefreshToken generate a new refresh token
func (p *PasetoToken) GenerateRefreshToken(userId uuid.UUID, duration time.Duration) (string, error) {
	return "", nil
}

// VerifyRefreshToken checks if a refresh token is valid or not
func (p *PasetoToken) VerifyRefreshToken(token string) (bool, error) {
	return false, nil
}

// NewPasetoToken create a new PasetoToken
func NewPasetoToken(secretKey string) (AccessToken, error) {
	if len(secretKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key size")
	}

	accessToken := &PasetoToken{
		paseto:    paseto.NewV2(),
		secretKey: []byte(secretKey),
	}
	return accessToken, nil
}
