package models

import (
	"time"

	"github.com/jtbry/CharlotteRoadReports/common"
)

type Incident struct {
	ID             string  `json:"eventNo"`
	TypeCode       string  `json:"typeCode"`
	TypeDesc       string  `json:"typeDescription"`
	TypeSubCode    string  `json:"typeSubCode"`
	TypeSubDesc    string  `json:"typeSubDescription"`
	Division       string  `json:"division"`
	XCord          int     `json:"xCoordinate"`
	YCord          int     `json:"yCoordinate"`
	Lat            float64 `json:"latitude"`
	Lon            float64 `json:"longitude"`
	Address        string  `json:"address"`
	DateTime       time.Time
	DateTimeString string `json:"eventDateTime,omitempty" gorm:"-"`
	IsActive       int
}

// Find all active incidents
func FindActiveIncidents() []Incident {
	actives := make([]Incident, 0)
	DB.Where("is_active = ?", 1).Order("date_time").Find(&actives)

	for i := 0; i < len(actives); i++ {
		actives[i].DateTimeString = common.UtcTimeToLocalString(actives[i].DateTime)
	}

	return actives
}
