package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTToken struct {
	secretKey []byte
}

type AccessTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateAccessToken create a new token for a specific user and durations
func (j *JWTToken) GenerateAccessToken(username string, duration time.Duration) (string, error) {
	claims := AccessTokenClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: duration.Milliseconds(),
			Issuer:    "pilaniya.auth.service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateRefreshToken generate a new refresh token
func (j *JWTToken) GenerateRefreshToken(username string, duration time.Duration) (string, error) {
	return "", nil
}

// VerifyAccessToken checks if a access token is valid or not
func (j *JWTToken) VerifyAccessToken(tokenString string) (bool, error) {

	claims :=AccessTokenClaims{}
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
