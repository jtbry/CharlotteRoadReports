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
