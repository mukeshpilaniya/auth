package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"sync"
)

type JWTToken struct {
	secretKey []byte
}

var (
	jwtTokenSingleton *JWTToken
	once sync.Once
)

// GenerateAccessToken create a new token for a specific user and durations
func (j *JWTToken) GenerateAccessToken(userId uuid.UUID) (string, error) {
	duration :=viper.GetDuration("ACCESS_TOKEN_DURATION")
	claims, err := NewToken(userId,duration)
	if err !=nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GenerateRefreshToken generate a new refresh token
func (j *JWTToken) GenerateRefreshToken(userId uuid.UUID) (string, error) {
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
	if jwtTokenSingleton==nil{
		once.Do(func() {
			jwtTokenSingleton = &JWTToken{
				secretKey: []byte(secretKey),
			}
		})
	}
	return jwtTokenSingleton, nil
}
