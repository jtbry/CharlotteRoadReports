package repository

import (
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"gorm.io/gorm"
)

type Storage interface {
	RunMigrations()
	FindActiveIncidents() []api.Incident
	FindIncidentById(eventNo string) api.Incident
	FindIncidentsWithFilter(filter api.IncidentFilter) []api.Incident
}

type storage struct {
	db *gorm.DB
}

// Create a new storage object from gorm.DB
func NewStorage(db *gorm.DB) Storage {
	return &storage{db: db}
}

// Run all required migrations for this storage
func (s *storage) RunMigrations() {
	s.db.AutoMigrate(&api.Incident{})
}

// Find all active incidents
func (s *storage) FindActiveIncidents() []api.Incident {
	actives := make([]api.Incident, 0)
	s.db.Where("is_active = ?", 1).Order("date_time").Find(&actives)
	return actives
}

// Find an incident by it's eventNo (ID)
func (s *storage) FindIncidentById(eventNo string) api.Incident {
	var incident api.Incident
	s.db.Where("id = ?", eventNo).Limit(1).Find(&incident)
	return incident
}

// Find all incidents that match the given filters
func (s *storage) FindIncidentsWithFilter(filter api.IncidentFilter) []api.Incident {
	query := s.db.Where("date_time >= ? AND date_time <= ? AND is_active = ?", filter.DateRangeStart, filter.DateRangeEnd, filter.ActivesOnly)
	results := make([]api.Incident, 0)
	query.Find(&results)
	return results
}
