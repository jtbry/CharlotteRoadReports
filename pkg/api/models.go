package api

import "time"

type Incident struct {
	// EventNo / ID for incident
	ID string
	// DateTimestamp the incident started
	StartTimestamp time.Time
	// DateTimestamp the incident ended
	EndTimestamp time.Time
	// Short version of TypeDesc
	TypeCode string
	// Description
	TypeDesc string
	// Short version of sub desc
	SubCode string
	// How the incident was logged / reported
	SubDesc string
	// CMPD divison
	Division string
	// Latitude
	Latitude float64
	// Longitude
	Longitude float64
	// Address
	Address string
	// Active status of the incident
	Active bool
}

type IncidentFilterRequest struct {
	// Start date for the date range
	DateRangeStart time.Time
	// End date for the date range
	DateRangeEnd time.Time
	// Whether or not Active must be true
	ActivesOnly bool
	// Address to use LIKE statement on
	AddressSearch string
}
