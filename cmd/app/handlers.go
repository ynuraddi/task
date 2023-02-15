package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"task/internal/app/user"
	"task/pkg/helpers"
	"task/pkg/validator"

	"github.com/go-chi/chi/v5"
)

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &user.UserModel{}

	if err := helpers.ReadJSON(w, r, &u); err != nil {
		log.Println(err)
		app.BadRequest(w, r, err)
		return
	}

	v := validator.New()
	v.Check(u.Email != "", "email", "must be provided")
	v.Check(validator.Matches(u.Email, validator.EmailRX), "email", "must be correct")
	v.Check(u.Password != "", "password", "must be provided")
	if !v.Valid() {
		app.BadRequest(w, r, errors.New(fmt.Sprint(v.Errors)))
		return
	}

	// kill routine after ctxTimeout
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	salt, err := http.Post(app.connect.saltURL, "", nil)
	if err != nil {
		log.Println(err)
		app.InternalServerError(w, r)
		return
	}

	dec := json.NewDecoder(salt.Body)
	if err := dec.Decode(u); err != nil {
		log.Println(err)
		app.InternalServerError(w, r)
		return
	}

	err = app.service.user.Create(ctx, u)
	switch {
	case errors.Is(err, user.ErrDuplicateEmail):
		log.Println(err)
		app.BadRequest(w, r, err)
	case err != nil:
		log.Println(err)
		app.InternalServerError(w, r)
	default:
		helpers.WriteJSON(w, http.StatusCreated, helpers.Envelope{"user": "created"})
	}
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	// kill routine after ctxTimeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := app.service.user.FindByEmail(ctx, email)
	if err != nil {
		log.Println(err)
		app.NotFound(w, r)
		return
	}

	data := helpers.Envelope{
		"email":    user.Email,
		"salt":     user.Salt,
		"password": user.Password,
	}

	if err := helpers.WriteJSON(w, http.StatusOK, data); err != nil {
		log.Println(err)
		app.InternalServerError(w, r)
		return
	}
}
