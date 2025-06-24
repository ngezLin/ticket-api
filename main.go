package main

import (
	"log"
	"os"
	"ticket-api/config"

	"ticket-api/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	db := config.SetupDatabase()
	router.SetupRoutes(r, db)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}