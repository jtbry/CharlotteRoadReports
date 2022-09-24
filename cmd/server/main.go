package main

import (
	"os"

	"github.com/jtbry/CharlotteRoadReports/internal/app"
	"github.com/jtbry/CharlotteRoadReports/pkg/repository"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006 Jan 06 15:04 MST",
		FullTimestamp:   true,
	})

	if os.Getenv("ENV") == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	if err := run(); err != nil {
		log.WithError(err).Error("Startup Error")
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
