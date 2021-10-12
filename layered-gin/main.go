// layered-gin/main.go

package main

import (
	HandlerCharacter "layered-gin/delivery/character"
	"layered-gin/driver"
	character "layered-gin/repository/character"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	// Load .env config file	
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env config file")
	}

	// Database connection
	conf := driver.DBConfig{
		Host: os.Getenv("SQL_HOST"),
		User: os.Getenv("SQL_USER"),
		Pass: os.Getenv("SQL_PASS"),
		Port: os.Getenv("SQL_PORT"),
		Name: os.Getenv("SQL_NAME"),
	}

	db, err := driver.InitDB(conf)
	if err != nil {
		log.Fatalf("unable to connect to sql")
	}

	// Router setup to handle requests
	repo := character.New(db)
	handler := HandlerCharacter.New(repo)

	router := gin.Default()

	api := router.Group("/api") 
	{
		api.GET("/characters/find/all", handler.GetAll)
		api.GET("/characters/find/:id", handler.GetByID)
		api.POST("/characters/create", handler.Create)
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("unable to start router")
	}
}