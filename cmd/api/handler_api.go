package main

import (
	"fmt"
	"github.com/mukeshpilaniya/auth/internal/models"
	"github.com/mukeshpilaniya/auth/internal/token"
	"github.com/mukeshpilaniya/auth/internal/util"
	"github.com/spf13/viper"
	"net/http"
)

// Payload is a type for sending custom message
type Payload struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

// getUserByID retrieve a user id of valid user
func (app *application) getUserByID(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err :=util.ReadJSON(w,r,&u)
	fmt.Println(u)
	if err != nil {
		app.errorLogger.Println(err)
		util.BadRequest(w,r,err)
		return
	}
	user, err :=app.DB.GetUserByID(u.ID)
	if err !=nil {
		app.errorLogger.Println(err)
		util.BadRequest(w,r,err)
		return
	}
	util.WriteJSON(w, http.StatusOK, &user)
}

// saveUser method save user into database
func (app *application) saveUser(w http.ResponseWriter, r *http.Request) {
 	 var u models.User

	 err :=util.ReadJSON(w,r,&u)
	 if err != nil {
		 app.errorLogger.Println(err)
		 util.BadRequest(w,r,err)
		 return
	 }
	 // check if the email is already present or not
	 _, err =app.DB.GetUserByEmail(u.Email)
	 if err == nil {
		 app.infoLogger.Println("user with email id already present")
		 var p Payload
		 p.Error = true
		 p.Message ="user with email id already present"
		 util.WriteJSON(w,http.StatusNotAcceptable,&p)
		 return
	 }
	 u, err =app.DB.SaveUser(u)
	if err !=nil {
		app.errorLogger.Println(err)
		util.BadRequest(w,r,err)
		return
	}
	util.WriteJSON(w, http.StatusOK, &u)
}

// generateAccessToken generate access token for a valid user
func (app *application) generateAccessToken(w http.ResponseWriter, r *http.Request) {
	var userCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := util.ReadJSON(w, r, &userCredentials)
	if err != nil {
		app.errorLogger.Println(err)
		util.BadRequest(w, r, err)
		return
	}

	// get the user by email
	u, err := app.DB.GetUserByEmail(userCredentials.Email)
	if err != nil {
		app.errorLogger.Println(err)
		util.InvalidCredentials(w)
		return
	}

	//check for password
	if string(u.Password) != string(userCredentials.Password) {
		util.InvalidCredentials(w)
		return
	}

	// generate token
	tokenGenerator, err := token.NewJWTToken(viper.GetString("TOKEN_SECRET_KEY"))
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	token, err := tokenGenerator.GenerateAccessToken(u.ID)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	// send response
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	var p Payload
	p.Error = false
	p.Message = "token generated"
	p.Token = token
	util.WriteJSON(w, http.StatusOK, &p)
}

// generateRefreshToken generate refresh token from a valid access token
func (app *application) generateRefreshToken(w http.ResponseWriter, r *http.Request ){

}