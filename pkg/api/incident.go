package api

import "time"

// Contains the required methods for IncidentService
type IncidentService interface {
	FindActiveIncidents() []Incident
}

// Allow storage interaction without being aware of the implementation
type IncidentRepository interface {
	FindActiveIncidents() []Incident
}

type incidentService struct {
	storage IncidentRepository
}

// Create a new IncidentService with the given repository
func NewIncidentService(repo IncidentRepository) IncidentService {
	return &incidentService{storage: repo}
}

// Find all active incidents
func (svc *incidentService) FindActiveIncidents() []Incident {
	actives := svc.storage.FindActiveIncidents()
	for i := 0; i < len(actives); i++ {
		actives[i].DateTimeString = utcToLocalString(actives[i].DateTime)
		actives[i].EndDateTimeString = utcToLocalString(*actives[i].EndDateTime)
		// N/A looks nicer than an empty address
		if actives[i].Address == "" {
			actives[i].Address = "N/A"
		}
	}
	return actives
}

// Convert UTC from storage to human readable Eastern Time
func utcToLocalString(t time.Time) string {
	location, _ := time.LoadLocation("America/New_York")
	return t.In(location).Format("1/2 - 3:04 PM")
}
