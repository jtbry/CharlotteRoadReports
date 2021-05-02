package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// todo: conform file to new structure, call a create incident function from incident service
// Poll the CMPD Public Data API every 3 minutes for traffic incidents
func beginPolling(db *gorm.DB) {
	for {
		// Execute in a goroutine to prevent it from blocking timer
		go func() {
			body, err := fetchCmpdData()
			if err != nil {
				fmt.Println(err)
			} else {
				processCmpdData(body, db)
			}
		}()

		// CMPD data updates once every three minutes
		<-time.After(3 * time.Minute)
	}
}

// Make GET request to CMPD API and process result
func fetchCmpdData() ([]byte, error) {
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

	return body, nil
}

// Process data fetched from the CMPD API
func processCmpdData(data []byte, db *gorm.DB) {
	// Unmarshal data
	activeIncidents := make([]api.Incident, 0)
	err := json.Unmarshal(data, &activeIncidents)
	if err != nil {
		fmt.Println(err)
	}

	// Fix incident values, get active incident IDs
	activeIDs := make([]string, len(activeIncidents))
	for i := 0; i < len(activeIncidents); i++ {
		activeIDs[i] = activeIncidents[i].ID
		// todo: could this be moved to a GORM hook? (BeforeCreate)
		activeIncidents[i].DateTime, err = parseIso8601Local(activeIncidents[i].DateTimeString)
		if err != nil {
			fmt.Printf("Unable to parse %s to time.Time\n", activeIncidents[i].DateTimeString)
		}
		activeIncidents[i].IsActive = 1
	}

	// Upsert to database
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&activeIncidents)

	// Update which incidents are active
	// WHERE is_active = 1 AND id NOT IN activeIds - UPDATE is_active = 0 AND end_date_time = NOW
	// EndDateTime has to be a pointer so that can be a null value in the database
	now := time.Now()
	db.Table("incidents").Where("is_active = 1").Not(map[string]interface{}{"id": activeIDs}).Select("is_active", "end_date_time").Updates(api.Incident{IsActive: 0, EndDateTime: &now})
}

// Parse an ISO8601 Local (Eastern Time) to time.Time
func parseIso8601Local(str string) (time.Time, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation("2006-01-02T15:04:05", str, location)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
