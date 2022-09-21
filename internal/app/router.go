package app

import (
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

func (s *Server) createApiRoutes() {
	// API Routes
	api := s.echo.Group("/api")
	api.GET("/incidents/active", s.getActiveIncidents)
	api.GET("/incidents/:eventNo", s.getIncidentById)
	api.GET("/incidents/search", s.searchIncidentsWithFilter)

	// SPA Routes if ServeClient is true
	if s.config.ServeClient {
		s.echo.Logger.Info("Serving SPA from web/build")
		s.echo.GET("/*", func(ctx echo.Context) error {
			url := path.Join("./web/build", ctx.Request().RequestURI)

			if _, err := os.Stat(url); os.IsNotExist(err) {
				return ctx.File("./web/build/index.html")
			} else {
				return ctx.File(url)
			}
		})
	}
}
