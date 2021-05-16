package api

import "time"

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

type IncidentFilter struct {
	// Start date for the date range
	DateRangeStart time.Time
	// End date for the date range
	DateRangeEnd time.Time
	// Whether or not IsActive must be true
	ActivesOnly int
}
