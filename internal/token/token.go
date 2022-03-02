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

// Token is a datatype of type token
type Token struct {
	// TokenID is a unique ID of JSON token
	TokenID uuid.UUID `json:"id,omitempty""`
	// UserID uniquely identify a user
	UserID uuid.UUID `json:"user_id"`
	// Issuer identifies the entity which issued the token.
	Issuer string `json:"iss,omitempty"`
	// IssuedAt is the time at which the token was issued.
	IssuedAt int64 `json:"issued_at"`
	// ExpiredAt is a time on or after which the token must not be accepted for processing.
	ExpiredAt int64 `json:"expired_at"`
	// claims is used for adding custom claims
	Claims map[string]string `json:"claims"`
}

// NewToken create a new Token for a specific username and duration
func NewToken(userId uuid.UUID, duration time.Duration) (*Token, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	token := &Token{
		TokenID:   tokenId,
		UserID:    userId,
		Issuer:    "pilaniya.auth.service",
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(duration).Unix(),
		Claims: make(map[string]string),
	}
	return token, nil
}

// Valid func check if the token is expired or not
func (token *Token) Valid() error {
	if token.ExpiredAt <= time.Now().Unix(){
		return ErrExpiredToken
	}
	return nil
}
