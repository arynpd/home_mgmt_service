package main

import (
	"log"

	"github.com/arynpd/home-mgmt-service/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := &db.Db{}
	err = db.Init()
	if err != nil {
		log.Fatalf("Error connecting to database: %s\n", err.Error())
	}
	defer db.Close()
}
