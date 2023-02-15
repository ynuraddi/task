package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"task/pkg/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var (
		host string
		port string
	)

	flag.StringVar(&host, "salt_host", "localhost", "host of salt-generator")
	flag.StringVar(&port, "salt_port", os.Getenv("salt_port"), "port of salt-generator")

	flag.Parse()

	r := route()

	log.Printf("Listen salt-generator %s:%s\n", host, port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func route() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	r.Post("/generate-salt", generateSalt)

	r.NotFound(NotFound)
	r.MethodNotAllowed(MethodNotAllowed)

	return r
}

func generateSalt(w http.ResponseWriter, r *http.Request) {
	salt := generateRandomSalt(saltSize)

	if err := helpers.WriteJSON(w, http.StatusCreated, helpers.Envelope{"salt": string(salt)}); err != nil {
		InternalServerError(w, r)
		return
	}
}
