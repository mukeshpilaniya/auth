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

// Payload contains payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload create a new Payload for a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid func check if the token is expired or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
