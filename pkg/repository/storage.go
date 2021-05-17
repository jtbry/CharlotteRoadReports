package repository

import (
	"time"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Storage interface {
	RunMigrations()
	FindActiveIncidents() []api.Incident
	FindIncidentById(eventNo string) api.Incident
	FindIncidentsWithFilter(filter api.IncidentFilter) []api.Incident
	UpdateActiveIncidents(actives []string)
	UpsertIncidentArray(incidents []api.Incident)
}

type pgsql struct {
	db *gorm.DB
}

// Create a new storage object from gorm.DB
func NewStorage(db *gorm.DB) Storage {
	return &pgsql{db: db}
}

// Run all required migrations for this storage
func (s *pgsql) RunMigrations() {
	s.db.AutoMigrate(&api.Incident{})
}

// Find all active incidents
func (s *pgsql) FindActiveIncidents() []api.Incident {
	actives := make([]api.Incident, 0)
	s.db.Where("active = ?", true).Order("start_timestamp").Find(&actives)
	return actives
}

// Find an incident by it's eventNo (ID)
func (s *pgsql) FindIncidentById(eventNo string) api.Incident {
	var incident api.Incident
	s.db.First(&incident, "id = ?", eventNo)
	return incident
}

// Find all incidents that match the given filters
func (s *pgsql) FindIncidentsWithFilter(filter api.IncidentFilter) []api.Incident {
	query := s.db.Where("start_timestamp >= ? AND start_timestamp <= ? AND active = ?", filter.DateRangeStart, filter.DateRangeEnd, filter.ActivesOnly)
	results := make([]api.Incident, 0)
	query.Order("start_timestamp").Find(&results)
	return results
}

// Update which incidents are active given an array of active IDs
func (s *pgsql) UpdateActiveIncidents(actives []string) {
	s.db.Table("incidents").Where("active = true").Not(map[string]interface{}{"id": actives}).Updates(map[string]interface{}{
		"active":        false,
		"end_timestamp": time.Now(),
	})
}

// Upsert a list of incidents updating the incident on conflict
func (s *pgsql) UpsertIncidentArray(incidents []api.Incident) {
	s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&incidents)
}
