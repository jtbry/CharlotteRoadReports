package app

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) setMiddleware() {
	if os.Getenv("ENV") == "development" {
		// Set development only middleware here
		s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${method}\t${path}\t${status}\t${latency_human}\n",
		}))
	}

	// Set shared middleware here
	s.echo.Use(middleware.Recover())
}
