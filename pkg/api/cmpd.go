package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

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

func cmpdIncidentToIncident(inc cmpdIncident) Incident {
	// Replace missing address with N/A
	if inc.Address == "" {
		inc.Address = "N/A"
	}

	// Load America/New_York time zone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.WithError(err).Warn("Failed to load America/New_York time zone")
		loc = time.UTC
	}

	// Convert string to time.Time
	startTimestamp, err := time.ParseInLocation("2006-01-02T15:04:05", inc.EventDateTime, loc)
	if err != nil {
		log.Warnf("(%s) Failed Timestamp Parse", inc.EventNo, inc.EventDateTime)
		startTimestamp = time.Now()
	}

	return Incident{
		ID:             inc.EventNo,
		StartTimestamp: startTimestamp,
		TypeCode:       inc.TypeCode,
		TypeDesc:       inc.TypeDescription,
		SubCode:        inc.TypeSubCode,
		SubDesc:        inc.TypeSubDescription,
		Division:       inc.Division,
		Latitude:       inc.Latitude,
		Longitude:      inc.Longitude,
		Address:        inc.Address,
		Active:         true,
	}
}

func FetchCmpdActiveIncidents() ([]Incident, error) {
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

	cmpdIncidents := make([]cmpdIncident, 0)
	err = json.Unmarshal(body, &cmpdIncidents)
	if err != nil {
		return nil, err
	}
	var incidents []Incident
	for i := range cmpdIncidents {
		incidents = append(incidents, cmpdIncidentToIncident(cmpdIncidents[i]))
	}

	return incidents, nil
}
