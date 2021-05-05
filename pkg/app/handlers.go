package app

import (
	"net/http"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleIncidentsActive() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, s.incidentService.FindActiveIncidents())
	}
}

func (s *Server) handleIncidentById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		eventNo := ctx.Param("eventNo")
		incident := s.incidentService.FindIncidentById(eventNo)
		if (incident == api.Incident{}) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": eventNo + " not found"})
		}
		return ctx.JSON(http.StatusOK, incident)
	}
}
