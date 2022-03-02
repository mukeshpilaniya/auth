package main

import (
	"fmt"
	"github.com/mukeshpilaniya/auth/internal/token"
	"github.com/mukeshpilaniya/auth/internal/util"
	"net/http"
	"time"
)

func (app *application) getUserByID(w http.ResponseWriter, r *http.Request) {

}

func (app *application) saveUser(w http.ResponseWriter, r *http.Request) {

}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
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
	tokenGenerator, err := token.NewPasetoToken("12348765123487651234876512348765")
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	token, err := tokenGenerator.GenerateAccessToken(u.ID, 5*time.Second)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	// send response
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Token   string `json:"token"`
	}
	payload.Error = false
	payload.Message = "token generated"
	payload.Token = token
	util.WriteJSON(w, http.StatusOK, &payload)
}
