package main

import (
	"net/http"

	"task/pkg/helpers"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 404, helpers.Envelope{"error": "route does not exist"})
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 405, helpers.Envelope{"error": "method is not valid"})
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, 500, helpers.Envelope{"error": "internal server error"})
}
