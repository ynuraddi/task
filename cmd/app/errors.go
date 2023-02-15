package main

import (
	"net/http"

	"task/pkg/helpers"
)

func (app *application) NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 404, helpers.Envelope{"error": "route does not exist"})
}

func (app *application) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	helpers.WriteJSON(w, 400, helpers.Envelope{"error": "bad request: " + err.Error()})
}

func (app *application) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 405, helpers.Envelope{"error": "method is not valid"})
}

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 500, helpers.Envelope{"error": "internal server error"})
}

func (app *application) DuplicateEmails(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 400, helpers.Envelope{"error": "email is already exist"})
}
