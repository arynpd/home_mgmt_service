package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arynpd/home-mgmt-service/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	controller := &controller.Controller{}
	err = controller.Init()
	if err != nil {
		log.Fatalf("Error connecting to database: %s\n", err.Error())
	}
	defer controller.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthcheck", controller.Healthcheck)

	http.ListenAndServe(os.Getenv("PORT"), router)
}
