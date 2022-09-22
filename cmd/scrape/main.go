package main

import (
	"os"
	"time"

	"github.com/jtbry/CharlotteRoadReports/internal/app"
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
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

	if config.ScheduledScraping {
		log.Info("Starting scheduled scraping")
		for {
			err = updateIncidentDatabase(pgsql)
			if err != nil {
				log.WithError(err).Error("updateIncidentDatabase failed")
			}

			<-time.After(time.Minute * 3)
		}

	} else {
		updateIncidentDatabase(pgsql)
	}
	return nil
}

func updateIncidentDatabase(repo api.IncidentRepository) error {
	incidents, err := api.FetchCmpdActiveIncidents()
	if err != nil {
		return err
	}

	activeEventIDs := make([]string, len(incidents))
	for i, incident := range incidents {
		activeEventIDs[i] = incident.ID
	}

	if len(activeEventIDs) < 1 {
		return nil
	}

	repo.UpsertIncidentArray(incidents)
	repo.UpdateActiveIncidents(activeEventIDs)
	log.Infof("Updated %d incidents", len(incidents))
	return nil
}
