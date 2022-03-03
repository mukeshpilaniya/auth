package main

import (
	"fmt"
	"github.com/mukeshpilaniya/auth/internal/models"
	"github.com/mukeshpilaniya/auth/internal/token"
	"github.com/mukeshpilaniya/auth/internal/util"
	"net/http"
	"time"
)


type Payload struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

func (app *application) getUserByID(w http.ResponseWriter, r *http.Request) {

}

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
	tokenGenerator, err := token.NewJWTToken("12348765123487651234876512348765")
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
	var p Payload
	p.Error = false
	p.Message = "token generated"
	p.Token = token
	util.WriteJSON(w, http.StatusOK, &p)
}
