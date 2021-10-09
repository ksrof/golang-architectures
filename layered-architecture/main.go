// layered-architecture/main.go

package main

import (
	"layered-architecture/datastore/character"
	HandlerCharacter "layered-architecture/delivery/character"
	"layered-architecture/driver"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	
	// loading env variables
	err = godotenv.Load()
	if err != nil {
		log.Println("could not load the .env file")
	}

	conf := driver.MySQLConfig{
		Host: os.Getenv("SQL_HOST"),
		User: os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port: os.Getenv("SQL_PORT"),
		DB: os.Getenv("SQL_DB"),
	}
	
	db, err := driver.MySQLConnect(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}

	datastore := character.New(db)
	handler := HandlerCharacter.New(datastore)

	http.HandleFunc("/character", handler.Handler)
	log.Println(http.ListenAndServe(":8080", nil))
}