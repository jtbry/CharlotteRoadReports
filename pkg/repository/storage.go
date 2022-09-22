package repository

import (
	"time"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type storage struct {
	db *gorm.DB
}

// Create a new storage object from gorm.DB
func NewStorage(dsn string, shouldMigrate bool) (api.IncidentRepository, error) {
	logger := logger.New(
		log.StandardLogger(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger,
	})
	if err != nil {
		return nil, err
	}

	s := &storage{db: db}
	if shouldMigrate {
		err = runMigrations(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// Run all required migrations for this storage
func runMigrations(s *storage) error {
	err := s.db.AutoMigrate(&api.Incident{})
	if err != nil {
		return err
	}
	return nil
}

// Find all active incidents
func (s *storage) FindActiveIncidents() []api.Incident {
	// TODO: remove sorting and move to client side when needed
	actives := make([]api.Incident, 0)
	s.db.Where("active = ?", true).Find(&actives)
	return actives
}

// Find an incident by it's eventNo (ID)
func (s *storage) FindIncidentById(eventNo string) api.Incident {
	var incident api.Incident
	s.db.First(&incident, "id = ?", eventNo)
	return incident
}

// Find all incidents that match the given filters
func (s *storage) FilterIncidents(filter api.IncidentFilterRequest) []api.Incident {
	query := s.db.Where("start_timestamp >= ? AND start_timestamp <= ?", filter.DateRangeStart, filter.DateRangeEnd)
	if filter.ActivesOnly {
		query = query.Where("active = true")
	}
	if filter.AddressSearch != "" {
		query = query.Where("address LIKE ?", "%"+filter.AddressSearch+"%")
	}
	results := make([]api.Incident, 0)
	query.Find(&results)
	return results
}

// Update which incidents are active given an array of active IDs
func (s *storage) UpdateActiveIncidents(actives []string) {
	s.db.Table("incidents").Where("active = true").Not(map[string]interface{}{"id": actives}).Updates(map[string]interface{}{
		"active":        false,
		"end_timestamp": time.Now(),
	})
}

// Upsert a list of incidents updating the incident on conflict
func (s *storage) UpsertIncidentArray(incidents []api.Incident) {
	s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&incidents)
}
