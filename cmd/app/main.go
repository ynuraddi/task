package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"task/internal/app/user"
	"task/pkg/mongodb"
)

type config struct {
	db struct {
		dsn            string
		collectionName string
	}
	salt struct {
		host string
		port string
	}
	app struct {
		host string
		port string
	}
}

type application struct {
	config config
	connect
	service
}

type connect struct {
	saltURL string
}

type service struct {
	user interface {
		Create(ctx context.Context, user *user.UserModel) error
		FindByEmail(ctx context.Context, email string) (*user.UserModel, error)
	}
}

const ctxTimeout = 5 * time.Second

func main() {
	var config config

	flag.StringVar(&config.db.dsn, "db_dsn", os.Getenv("db_dsn"), "dsn for db")
	flag.StringVar(&config.db.collectionName, "db_collection_name", "users", "dsn for db")

	flag.StringVar(&config.salt.host, "salt_host", os.Getenv("salt_host"), "host of salt-generator")
	flag.StringVar(&config.salt.port, "salt_port", os.Getenv("salt_port"), "port of salt-generator")

	flag.StringVar(&config.app.host, "app_host", "localhost", "host of application")
	flag.StringVar(&config.app.port, "app_port", os.Getenv("app_port"), "port of application")

	flag.Parse()

	mongodb, err := mongodb.NewClient(config.db.dsn)
	if err != nil {
		panic(err)
	}

	userRepo := user.NewUserRepository(mongodb, config.db.collectionName)

	userServ := user.NewUserService(userRepo)

	app := &application{
		config: config,
		connect: connect{
			saltURL: fmt.Sprintf("http://%s:%s/generate-salt", config.salt.host, config.salt.port),
		},
		service: service{
			user: userServ,
		},
	}

	log.Printf("Listen user-creator localhost:%s\n", config.app.port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", config.app.port), app.routes()))
}
