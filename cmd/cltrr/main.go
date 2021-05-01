package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/jtbry/CharlotteRoadReports/pkg/app"
	"github.com/jtbry/CharlotteRoadReports/pkg/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup Error: %s\n", err)
		return
	}
}

func run() error {
	// Check for required environment variables
	if err := checkEnvironment(); err != nil {
		return err
	}

	// Connect database
	db, err := newDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	// Begin polling, for now this is integrated in the same process as the web server
	// If hosting limitations change this can be moved to it's own process as a cron job
	go beginPolling(db)

	// Create storage and Run required migrations
	storage := repository.NewStorage(db)
	storage.RunMigrations()

	// Create router
	router := gin.New()

	// Create services
	incidentService := api.NewIncidentService(storage)

	// Create and start server
	server := app.NewServer(router, incidentService)
	err = server.Run()
	if err != nil {
		return err
	}

	return nil
}

func checkEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		// .env isn't always required
		fmt.Println("Error loading .env file")
		fmt.Println(err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		// Assume development if no env is given
		os.Setenv("ENV", "development")
	}

	port := os.Getenv("PORT")
	if port == "" {
		// Port is always required
		return errors.New("$PORT must be set")
	}
	return nil
}

func newDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}