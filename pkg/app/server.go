package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
)

// Contains the params required for web server operations
type Server struct {
	router          *gin.Engine
	incidentService api.IncidentService
}

// Create a new server object
func NewServer(router *gin.Engine, incidentService api.IncidentService) *Server {
	return &Server{
		router:          router,
		incidentService: incidentService,
	}
}

// Setup and run the server
func (s *Server) Run() error {
	s.setMiddleware()
	s.setRoutes()
	err := s.router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return err
	}
	return nil
}
