package app

import (
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

func (s *Server) setRoutes() {
	// Serve API routes
	api := s.echo.Group("/api")
	api.GET("/incidents/active", s.handleIncidentsActive())
	api.GET("/incidents/:eventNo", s.handleIncidentById())

	// Serve react app
	// This will only serve the production react app if available
	// debugging the react app should be done using npm start
	s.echo.GET("/*", func(ctx echo.Context) error {
		url := path.Join("./frontend/build", ctx.Request().RequestURI)

		if _, err := os.Stat(url); os.IsNotExist(err) {
			return ctx.File("./frontend/build/index.html")
		} else {
			return ctx.File(url)
		}
	})
}
