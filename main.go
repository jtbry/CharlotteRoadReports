package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/models"
	"github.com/jtbry/CharlotteRoadReports/routes"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	// Set up environment
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	models.DatabaseInit()
	go BeginPolling()

	// Create web app.
	web := gin.New()
	web.Use(gin.Recovery())
	if env == "development" {
		web.Use(gin.Logger())
	} else {
		nrlicense := os.Getenv("NEW_RELIC_LICENSE_KEY")
		if nrlicense != "" {
			app, err := newrelic.NewApplication(
				newrelic.ConfigAppName("CharlotteRoadReport"),
				newrelic.ConfigLicense(nrlicense),
			)
			if err != nil {
				log.Printf("%s\nUnable to create New Relic app\n", err)
			} else {
				web.Use(nrgin.Middleware(app))
			}
		}
	}

	api := web.Group("/api")
	routes.RegisterIncidentsApi(api)
	web.Run(":" + port)
}
