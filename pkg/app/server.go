package app

import (
	"os"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/labstack/echo/v4"
)

// Contains the params required for web server operations
type Server struct {
	echo            *echo.Echo
	incidentService api.IncidentService
}

// Create a new server object
func NewServer(e *echo.Echo, incidentService api.IncidentService) *Server {
	return &Server{
		echo:            e,
		incidentService: incidentService,
	}
}

// Setup and run the server
func (s *Server) Run() error {
	s.setMiddleware()
	s.setRoutes()

	if os.Getenv("ENV") == "development" {
		s.echo.Debug = true
	}

	err := s.echo.Start(":" + os.Getenv("PORT"))
	if err != nil {
		return err
	}
	return nil
}
