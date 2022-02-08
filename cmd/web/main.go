package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/pkg/app"
	"github.com/jtbry/CharlotteRoadReports/pkg/repository"
	"github.com/labstack/echo/v4"
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

	// Create pgsql storage / db connection
	pgsql, err := repository.NewPgsqlStorage(os.Getenv("DATABASE_URL"), true)
	if err != nil {
		return err
	}

	// Create web server
	e := echo.New()

	// Create and start server
	server := app.NewServer(e, pgsql)
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
