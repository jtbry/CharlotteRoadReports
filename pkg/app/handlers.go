package app

import (
	"net/http"
	"strconv"
	"time"

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

func (s *Server) handleIncidentSearch() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		drs, err := time.Parse(time.RFC3339, ctx.QueryParam("dateRangeStart"))
		if err != nil {
			return err
		}
		dre, err := time.Parse(time.RFC3339, ctx.QueryParam("dateRangeEnd"))
		if err != nil {
			return err
		}
		ao, err := strconv.Atoi(ctx.QueryParam("activesOnly"))
		if err != nil {
			return err
		}

		filter := api.IncidentFilter{
			DateRangeStart: drs,
			DateRangeEnd:   dre,
			ActivesOnly:    ao,
		}
		return ctx.JSON(http.StatusOK, s.incidentService.FindIncidentsWithFilter(filter))
	}
}
