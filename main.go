package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/models"
	"github.com/jtbry/CharlotteRoadReports/routes"
)

func main() {
	// Set up environment
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	models.DatabaseInit()
	go BeginPolling()

	// Create web app.
	web := gin.Default()
	api := web.Group("/api")
	routes.RegisterIncidentsApi(api)
	web.Run(":" + port)
}
