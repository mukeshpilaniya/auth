package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mukeshpilaniya/auth/pkg/midware"
	"net/http"
)

func (app *application) routes() http.Handler{
	mux := chi.NewRouter()

	// Common middleware for all routes
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.URLFormat)

	// public routes
	mux.Post("/api/v1/generate_token", app.generateToken)
	mux.Post("/api/v1/generate_access_token",app.generateAccessTokenFromRefreshToken)
	mux.Post("/api/v1/create_user",app.createUser)

	// protected routes
	mux.Group(func (r chi.Router){
		r.Use(midware.JWTAuthMiddleware)
		r.Post("/api/v1/get_user", app.getUserByID)
	})
	return mux
}
