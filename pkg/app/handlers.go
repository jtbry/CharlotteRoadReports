package app

import (
	"net/http"
	"time"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleIncidentsActive() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, s.incidentRepo.FindActiveIncidents())
	}
}

func (s *Server) handleIncidentById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		eventNo := ctx.Param("eventNo")
		incident := s.incidentRepo.FindIncidentById(eventNo)
		if (incident == api.Incident{}) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"error": eventNo + " not found"})
		}
		return ctx.JSON(http.StatusOK, incident)
	}
}

func (s *Server) handleIncidentSearch() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		filter := api.IncidentFilter{}
		err := echo.QueryParamsBinder(ctx).
			Time("dateRangeStart", &filter.DateRangeStart, time.RFC3339).
			Time("dateRangeEnd", &filter.DateRangeEnd, time.RFC3339).
			Bool("activesOnly", &filter.ActivesOnly).
			String("addressSearch", &filter.AddressSearch).
			BindError()
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, s.incidentRepo.FindIncidentsWithFilter(filter))
	}
}
