package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/jtbry/CharlotteRoadReports/pkg/repository"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup Error: %s\n", err)
		return
	}
}

func run() error {
	// Load env
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("no DSN provided")
	}

	// Create pgsql storage / db connection
	pgsql, err := repository.NewPgsqlStorage(dsn, true)
	if err != nil {
		return err
	}

	// Only use in-app scheduling when needed
	shouldSchedule := os.Getenv("SCHEDULE") == "true"
	api.BeginPolling(pgsql, shouldSchedule)

	return nil
}
