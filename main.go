package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting the Rick and Morty Dex")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	port := ":" + os.Getenv("PORT")
	mode := os.Getenv("GIN_MODE")

	if mode != "debug" && mode != "release" && mode != "test" {
		mode = "debug"
	}

	gin.SetMode(mode)
	// Initialize the Gin router
	r := gin.Default()
	
	log.Println("Starting server on port", port)

	// Start the server
	r.Run(port)
}
