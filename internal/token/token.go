package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

// Different types of errors return by the VerifyToken func
var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token in invalid")
)

// Token contains Token data of the token
type Token struct {
	ID        uuid.UUID `json:"id"`
	username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewToken create a new Token for a specific username and duration
func NewToken(username string, duration time.Duration) (*Token, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	token := &Token{
		ID:        tokenId,
		username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return token, nil
}

// Valid func check if the token is expired or not
func (token *Token) Valid() error {
	if time.Now().After(token.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
