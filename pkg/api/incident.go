package api

// Allow storage interaction without being aware of the implementation
type IncidentRepository interface {
	FindActiveIncidents() []Incident
	FindIncidentById(eventNo string) Incident
	FindIncidentsWithFilter(filter IncidentFilter) []Incident
	UpdateActiveIncidents(actives []string)
	UpsertIncidentArray(incidents []Incident)
}

type incidentRepository struct {
	storage IncidentRepository
}

// Create a new IncidentRepository with the given repository/storage object
// The repository.Storage interface will be implemented as an incident repository
func NewIncidentRepo(repo IncidentRepository) IncidentRepository {
	return &incidentRepository{storage: repo}
}

// Find all active incidents
func (repo *incidentRepository) FindActiveIncidents() []Incident {
	actives := repo.storage.FindActiveIncidents()
	return actives
}

// Find an incident by eventNo
func (repo *incidentRepository) FindIncidentById(eventNo string) Incident {
	return repo.storage.FindIncidentById(eventNo)
}

// Find all incidents that match the given filters
func (repo *incidentRepository) FindIncidentsWithFilter(filter IncidentFilter) []Incident {
	return repo.storage.FindIncidentsWithFilter(filter)
}

// Update which incidents are active given an array of active IDs
func (repo *incidentRepository) UpdateActiveIncidents(actives []string) {
	repo.storage.UpdateActiveIncidents(actives)
}

// Upsert a list of incidents updating the incident on conflict
func (repo *incidentRepository) UpsertIncidentArray(incidents []Incident) {
	repo.storage.UpsertIncidentArray(incidents)
}
