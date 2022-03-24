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
	//jwtTokenRefreshSingleton *JWTToken
	once sync.Once
)

// GenerateAccessToken create a new token for a specific user and durations
func (j *JWTToken) GenerateAccessToken(userId uuid.UUID) (string, error) {
	duration := viper.GetDuration("ACCESS_TOKEN_DURATION")
	claims, err := NewToken(userId, duration)
	if err != nil {
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
	duration := viper.GetDuration("REFRESH_TOKEN_DURATION")
	claims, err := NewToken(userId, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyAccessToken checks if a access token is valid or not
func (j *JWTToken) VerifyAccessToken(tokenString string) (uuid.UUID, bool, error) {
	var id uuid.UUID
	claims := &Token{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return id, false, err
	}
	if !token.Valid {
		return id, false, jwt.ErrSignatureInvalid
	}
	id = claims.UserID
	return id, true, nil
}

// VerifyRefreshToken checks if a refresh token is valid or not
func (j *JWTToken) VerifyRefreshToken(tokenString string) (uuid.UUID, bool, error) {
	var id uuid.UUID
	claims := &Token{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return id, false, err
	}
	if !token.Valid {
		return id, false, jwt.ErrSignatureInvalid
	}
	id = claims.UserID
	return id, true, nil
}

//// NewJWTAccessToken return a singleton instance of type JWTToken
//func NewJWTAccessToken(secretKey string) (*JWTToken, error) {
//	if jwtTokenAccessSingleton==nil{
//		once.Do(func() {
//			jwtTokenAccessSingleton = &JWTToken{
//				secretKey: []byte(secretKey),
//			}
//		})
//	}
//	return jwtTokenAccessSingleton, nil
//}
//
//// NewJWTRefreshToken return a singleton instance of type JWTToken
//func NewJWTRefreshToken(secretKey string) (*JWTToken, error) {
//	if jwtTokenRefreshSingleton==nil{
//		once.Do(func() {
//			jwtTokenRefreshSingleton = &JWTToken{
//				secretKey: []byte(secretKey),
//			}
//		})
//	}
//	return jwtTokenRefreshSingleton, nil
//}

// NewJWTToken return a singleton instance of type JWTToken
func NewJWTToken(secretKey string) (*JWTToken, error) {
	jwtTokenSingleton = &JWTToken{
		secretKey: []byte(secretKey),
	}
	return jwtTokenSingleton, nil
}
