package main

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mukeshpilaniya/auth/internal/models"
	"github.com/mukeshpilaniya/auth/internal/token"
	"github.com/mukeshpilaniya/auth/internal/util"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

// Payload is a type for sending custom message
type Payload struct {
	Error        bool   `json:"error,omitempty"`
	Message      string `json:"message,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

var userCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrInvalidUser = errors.New("username or password are not correct")
)

// getUserByID retrieve a user id of valid user
func (app *application) getUserByID(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := util.ReadJSON(w, r, &u)
	if err != nil {
		app.errorLogger.Println(err)
		util.BadRequest(w, r, err)
		return
	}
	user, err := app.DB.GetUserByID(u.ID)
	if err != nil {
		app.errorLogger.Println(err)
		util.BadRequest(w, r, err)
		return
	}
	util.WriteJSON(w, http.StatusOK, &user)
}

// createUser save user into database
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := util.ReadJSON(w, r, &u)
	if err != nil {
		app.errorLogger.Println(err)
		util.BadRequest(w, r, err)
		return
	}
	// check if the email is already present or not
	_, err = app.DB.GetUserByEmail(u.Email)
	if err == nil {
		app.infoLogger.Println("user with email id already present")
		var p Payload
		p.Error = true
		p.Message = "user with email id already present"
		util.WriteJSON(w, http.StatusNotAcceptable, &p)
		return
	}
	u, err = app.DB.SaveUser(u)
	if err != nil {
		app.errorLogger.Println(err)
		util.BadRequest(w, r, err)
		return
	}
	util.WriteJSON(w, http.StatusOK, &u)
}

// validateUser check whether the user is valid or not
func (app *application) validateUser(w http.ResponseWriter, r *http.Request) (models.User, bool, error) {
	err := util.ReadJSON(w, r, &userCredentials)
	var u models.User
	if err != nil {
		util.BadRequest(w, r, err)
		return u, false, err
	}

	// get the user by email
	u, err = app.DB.GetUserByEmail(userCredentials.Email)
	if err != nil {
		util.InvalidCredentials(w)
		return u, false, err
	}

	//check for password
	if strings.ToLower(u.Password) != strings.ToLower(userCredentials.Password) {
		util.InvalidCredentials(w)
		return u, false, ErrInvalidUser
	}
	return u, true, nil
}

// generateAccessToken generate access token for a valid user
func (app *application) generateAccessToken(id uuid.UUID) (string, error) {
	// generate authentication token
	tokenGenerator, err := token.NewJWTToken(viper.GetString("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		return "", err
	}
	accessToken, err := tokenGenerator.GenerateAccessToken(id)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// generateRefreshToken generate refresh token for a valid user
func (app *application) generateRefreshToken(id uuid.UUID) (string, error) {
	// generate authentication token
	tokenGenerator, err := token.NewJWTToken(viper.GetString("REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		return "", err
	}
	refreshToken, err := tokenGenerator.GenerateRefreshToken(id)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

// generateToken generate access and refresh token for a valid user
func (app *application) generateToken(w http.ResponseWriter, r *http.Request) {
	u, isValid, err := app.validateUser(w, r)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	if !isValid {
		app.errorLogger.Println(ErrInvalidUser)
		return
	}

	accessTokenString, err := app.generateAccessToken(u.ID)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	refreshTokenString, err := app.generateRefreshToken(u.ID)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	// send response back
	var p Payload
	p.Error = false
	p.Message = "token generated"
	p.AccessToken = accessTokenString
	p.RefreshToken = refreshTokenString
	util.WriteJSON(w, http.StatusOK, &p)
}

// validateRefreshToken check whether a refresh token is valid or not
func (app *application) validateRefreshToken(w http.ResponseWriter, r *http.Request) (uuid.UUID,bool, error) {
	var validateToken struct{
		RefreshToken string `json:"refresh_token"`
	}
	var id uuid.UUID
	err := util.ReadJSON(w, r, &validateToken)
	if err != nil {
		return id,false, err
	}
	secretKey := viper.GetString("REFRESH_TOKEN_SECRET_KEY")
	jwtToken, err := token.NewJWTToken(secretKey)
	if err != nil {
		util.WriteJSON(w, http.StatusNotAcceptable, util.Payload{Message: "internal server error", Error: true}, nil)
		return id,false, err
	}
	id, isValid, err := jwtToken.VerifyRefreshToken(validateToken.RefreshToken)
	if err != nil {
		util.WriteJSON(w, http.StatusUnauthorized, util.Payload{Message: err.Error(), Error: true}, nil)
		return id,false, err
	}
	if isValid == false {
		util.WriteJSON(w, http.StatusUnauthorized, util.Payload{Message: "unauthorized", Error: true}, nil)
		return id,false, err
	}
	return id, true, nil
}

func (app *application) generateAccessTokenFromRefreshToken(w http.ResponseWriter, r *http.Request) {
	id, isValid, err := app.validateRefreshToken(w, r)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	if !isValid {
		app.errorLogger.Println(err)
		return
	}
	accessTokenString,err := app.generateAccessToken(id)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	// send response back
	var p Payload
	p.Error = false
	p.Message = "Access token generated"
	p.AccessToken = accessTokenString
	util.WriteJSON(w, http.StatusOK, &p)
}
