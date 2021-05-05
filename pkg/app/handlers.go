package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleIncidentsActive() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, s.incidentService.FindActiveIncidents())
	}
}
