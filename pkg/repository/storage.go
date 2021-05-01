package repository

import (
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"gorm.io/gorm"
)

type Storage interface {
	RunMigrations()
	FindActiveIncidents() []api.Incident
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
