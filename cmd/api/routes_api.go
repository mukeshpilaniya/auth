package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (app *application) routes() http.Handler{
	mux := chi.NewRouter()

	mux.Get("/api/user/{id}", app.getUserByID)
	//mux.Post("/api/login",app.login)
	mux.Post("/api/user",app.saveUser)
	mux.Post("/api/generate_access_token",app.generateAccessToken)
	mux.Post("/api/generate_refresh_token",app.generateRefreshToken)

	return mux
}
