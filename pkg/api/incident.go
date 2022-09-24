package api

type IncidentRepository interface {
	// Find all active incidents
	FindActiveIncidents() []Incident
	// Find an incident by eventNo
	FindIncidentById(eventNo string) Incident
	// Find all incidents that match the given filter
	FilterIncidents(filter IncidentFilterRequest) []Incident
	// Update which incidents are active given an array of active IDs
	UpdateActiveIncidents(incidents []Incident)
}
