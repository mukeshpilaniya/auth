package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mukeshpilaniya/auth/pkg/midware"
	"net/http"
)

func (app *application) routes() http.Handler{
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	// public routes
	//mux.Use(middleware.LoggingMiddleware)
	//mux.Post("/api/login",app.login)
	mux.Post("/api/generate_access_token",app.generateAccessToken)

	// protected routes
	mux.Group(func (r chi.Router){
		r.Use(midware.AuthMiddleware)
		r.Post("/api/user/", app.getUserByID)
		//r.Post("/api/user",app.saveUser)
		r.Post("/api/generate_refresh_token",app.generateRefreshToken)
	})
	return mux
}
