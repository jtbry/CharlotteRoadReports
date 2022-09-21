package app

import (
	"net/http"
	"time"

	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"github.com/labstack/echo/v4"
)

func (s *Server) getActiveIncidents(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, s.incidentRepo.FindActiveIncidents())
}

func (s *Server) getIncidentById(ctx echo.Context) error {
	eventNo := ctx.Param("eventNo")
	incident := s.incidentRepo.FindIncidentById(eventNo)
	if (incident == api.Incident{}) {
		return ctx.JSON(http.StatusNotFound, nil)
	}
	return ctx.JSON(http.StatusOK, incident)
}

func (s *Server) searchIncidentsWithFilter(ctx echo.Context) error {
	filter := api.IncidentFilterRequest{}
	err := echo.QueryParamsBinder(ctx).
		Time("dateRangeStart", &filter.DateRangeStart, time.RFC3339).
		Time("dateRangeEnd", &filter.DateRangeEnd, time.RFC3339).
		Bool("activesOnly", &filter.ActivesOnly).
		String("addressSearch", &filter.AddressSearch).
		BindError()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, s.incidentRepo.FilterIncidents(filter))
}
