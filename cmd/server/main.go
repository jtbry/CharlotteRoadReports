package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jtbry/CharlotteRoadReports/internal/app"
	"github.com/jtbry/CharlotteRoadReports/pkg/repository"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "[%s] Startup Error: %s\n", time.Now().Format(time.ANSIC), err)
		return
	}
}

func run() error {
	config, err := app.LoadConfig()
	if err != nil {
		return err
	}

	pgsql, err := repository.NewStorage(config.DatabaseURL, true)
	if err != nil {
		return err
	}

	server := app.NewServer(config, pgsql)
	return server.Run()
}
