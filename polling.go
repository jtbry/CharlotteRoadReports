package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jtbry/CharlotteRoadReports/common"
	"github.com/jtbry/CharlotteRoadReports/models"
	"gorm.io/gorm/clause"
)

// Process data fetched from the CMPD API
func processCmpdData(data []byte) {
	// Unmarshal data
	activeIncidents := make([]models.Incident, 0)
	err := json.Unmarshal(data, &activeIncidents)
	if err != nil {
		fmt.Println(err)
	}

	// Fix incident values, get active incident IDs
	activeIDs := make([]string, len(activeIncidents))
	for i := 0; i < len(activeIncidents); i++ {
		activeIDs[i] = activeIncidents[i].ID
		// todo: could this be moved to a GORM hook? (BeforeCreate)
		activeIncidents[i].DateTime, err = common.ParseIso8601Local(activeIncidents[i].DateTimeString)
		if err != nil {
			fmt.Printf("Unable to parse %s to time.Time\n", activeIncidents[i].DateTimeString)
		}
		activeIncidents[i].IsActive = 1
	}

	// Upsert to database
	models.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&activeIncidents)

	// Update which incidents are active
	models.DB.Table("incidents").Not(map[string]interface{}{"id": activeIDs}).Update("is_active", 0)
}

// Make GET request to CMPD API and process result
func fetchCmpdData() {
	url := "https://cmpdinfo.charlottenc.gov/api/v2.1/traffic"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	processCmpdData(body)
}

// Poll the CMPD Public Data API every 3 minutes for traffic incidents
func BeginPolling() {
	for {
		go fetchCmpdData()
		// CMPD data updates once every three minutes
		<-time.After(3 * time.Minute)
	}
}
