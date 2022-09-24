package app

import (
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Contains the params required for web server operations
type Server struct {
	echo         *echo.Echo
	incidentRepo api.IncidentRepository
	config       AppConfig
}

// Create a new server object
func NewServer(config AppConfig, incidentRepo api.IncidentRepository) *Server {
	return &Server{
		echo:         echo.New(),
		incidentRepo: incidentRepo,
		config:       config,
	}
}

// Setup and run the http server using echo
func (s *Server) Run() error {
	s.createApiRoutes()

	if s.config.Env == "development" {
		s.echo.Debug = true
		s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${method}\t${path}\t${status}\t${latency_human}\n",
		}))
	}
	s.echo.Use(middleware.Recover())

	err := s.echo.Start(":" + s.config.Port)
	if err != nil {
		return err
	}
	return nil
}
