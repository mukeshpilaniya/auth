package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type JWTToken struct {
	secretKey []byte
}

// GenerateAccessToken create a new token for a specific user and durations
func (j *JWTToken) GenerateAccessToken(userId uuid.UUID, duration time.Duration) (string, error) {
	tokenId, err := uuid.NewRandom()
	if err !=nil {
		return "",err
	}
	claims := &Token{
		TokenID: tokenId,
		UserID: userId,
		ExpiredAt: duration.Milliseconds(),
		IssuedAt: time.Now().Unix(),
		Issuer: "pilaniya.auth.service",
		Claims: make(map[string]string),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateRefreshToken generate a new refresh token
func (j *JWTToken) GenerateRefreshToken(userId uuid.UUID, duration time.Duration) (string, error) {
	return "", nil
}

// VerifyAccessToken checks if a access token is valid or not
func (j *JWTToken) VerifyAccessToken(tokenString string) (bool, error) {

	claims :=&Token{}
	token, err :=jwt.ParseWithClaims(tokenString,claims, func(tkn *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err !=nil {
		return false, err
	}
	if !token.Valid{
		return false, jwt.ErrSignatureInvalid
	}
	return true, nil
}

// VerifyRefreshToken checks if a refresh token is valid or not
func (j *JWTToken) VerifyRefreshToken(tokenString string) (*Token, error) {
	return &Token{}, nil
}

func NewJWTToken(secretKey string) (*JWTToken, error) {
	j := &JWTToken{
		secretKey: []byte(secretKey),
	}
	return j, nil
}
