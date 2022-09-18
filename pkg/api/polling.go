package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Begin polling data sources for incidents
func BeginPolling(repo IncidentRepository, shouldSchedule bool) {
	fetch := func() {
		incidents, err := pollCmpd()
		if err != nil {
			fmt.Println(err)
			return
		}

		activesLen := len(incidents)
		activeIDs := make([]string, activesLen)
		for i := 0; i < activesLen; i++ {
			activeIDs[i] = incidents[i].ID
		}

		if activesLen <= 0 {
			// Don't execute if there are no active events
			// Prevents a verbose error from GORM
			return
		}
		repo.UpsertIncidentArray(incidents)
		repo.UpdateActiveIncidents(activeIDs)
		fmt.Printf("Polled %d active incidents at %s\n", activesLen, time.Now().UTC().Format(time.ANSIC))
	}

	if shouldSchedule {
		for {
			// Collect data in a separate goroutine to prevent blocking
			go fetch()

			// CMPD data only updates once every 3min
			<-time.After(3 * time.Minute)
		}
	} else {
		fetch()
	}
}

// Struct to unmarshal CMPD data into before converting to Incident
type cmpdIncident struct {
	EventNo             string  `json:"eventNo"`
	EventDateTime       string  `json:"eventDateTime"`
	AddedDateTimeString string  `json:"addedDateTimeString"`
	TypeCode            string  `json:"typeCode"`
	TypeDescription     string  `json:"typeDescription"`
	TypeSubCode         string  `json:"typeSubCode"`
	TypeSubDescription  string  `json:"typeSubDescription"`
	Division            string  `json:"division"`
	XCoordinate         int     `json:"xCoordinate"`
	YCoordinate         int     `json:"yCoordinate"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Address             string  `json:"address"`
}

// Poll the CMPD data sources for incidents
func pollCmpd() ([]Incident, error) {
	cmpdIncidents, err := fetchCmpdIncidents()
	if err != nil {
		return nil, err
	}

	// Do any processing and convert to an Incident
	var incidents []Incident
	for i := 0; i < len(cmpdIncidents); i++ {
		// Replace missing addresses with N/A
		if cmpdIncidents[i].Address == "" {
			cmpdIncidents[i].Address = "N/A"
		}

		// Convert string to time.Time
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			fmt.Println(err)
			continue
		}

		startTimestamp, err := time.ParseInLocation("2006-01-02T15:04:05", cmpdIncidents[i].EventDateTime, loc)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Append cmpd incident as an Incident
		incidents = append(incidents, Incident{
			ID:             cmpdIncidents[i].EventNo,
			StartTimestamp: startTimestamp,
			TypeCode:       cmpdIncidents[i].TypeCode,
			TypeDesc:       cmpdIncidents[i].TypeDescription,
			SubCode:        cmpdIncidents[i].TypeSubCode,
			SubDesc:        cmpdIncidents[i].TypeSubDescription,
			Division:       cmpdIncidents[i].Division,
			Latitude:       cmpdIncidents[i].Latitude,
			Longitude:      cmpdIncidents[i].Longitude,
			Address:        cmpdIncidents[i].Address,
			Active:         true,
		})
	}
	return incidents, nil
}

// Make HTTP request to CMPD api
func fetchCmpdIncidents() ([]cmpdIncident, error) {
	url := "https://cmpdinfo.charlottenc.gov/api/v2.1/traffic"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	incidents := make([]cmpdIncident, 0)
	err = json.Unmarshal(body, &incidents)
	if err != nil {
		return nil, err
	}
	return incidents, nil
}
