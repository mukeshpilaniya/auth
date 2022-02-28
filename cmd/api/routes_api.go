package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (app *application) routes() http.Handler{
	mux := chi.NewRouter()

	mux.Get("/api/user/{id}", app.getUserByID)
	mux.Post("/api/login",app.login)
	mux.Post("/api/user",app.saveUser)

	return mux
}
