package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/jtbry/CharlotteRoadReports/pkg/app"
	"github.com/jtbry/CharlotteRoadReports/pkg/repository"
	"github.com/labstack/echo/v4"
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

	// Create storage and Run required migrations
	storage := repository.NewStorage(db)
	storage.RunMigrations()

	// Create repositories
	incidentRepo := api.NewIncidentRepo(storage)

	// Create web server
	e := echo.New()

	// Begin polling, for now this is integrated in the same process as the web server
	// If hosting limitations change this can be moved to it's own process as a cron job
	go beginPolling(incidentRepo)

	// Create and start server
	server := app.NewServer(e, incidentRepo)
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
