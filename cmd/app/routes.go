package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.AllowContentType("application/json"))

	r.Post("/create-user", app.CreateUser)
	r.Get("/get-user/{email}", app.GetUser)

	r.NotFound(app.NotFound)
	r.MethodNotAllowed(app.MethodNotAllowed)

	return r
}
