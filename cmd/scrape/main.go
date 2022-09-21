package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jtbry/CharlotteRoadReports/internal/app"
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
	config, err := app.LoadConfig()
	if err != nil {
		return err
	}

	pgsql, err := repository.NewStorage(config.DatabaseURL, true)
	if err != nil {
		return err
	}

	if config.ScheduledScraping {
		for {
			err = updateIncidentDatabase(pgsql)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fetch Error: %s\n", err)
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
	fmt.Printf("[%s] Fetched %d active incidents\n", time.Now().UTC().Format(time.ANSIC), len(incidents))
	return nil
}
