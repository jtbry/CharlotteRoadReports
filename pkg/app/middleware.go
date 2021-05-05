package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func newRelicMiddleware() (echo.MiddlewareFunc, error) {
	license := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if license == "" {
		return nil, errors.New("missing $NEW_RELIC_LICENSE_KEY")
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("CharlotteRoadReport"),
		newrelic.ConfigLicense(license),
	)
	if err != nil {
		return nil, err
	}
	return nrecho.Middleware(app), nil
}

func (s *Server) setMiddleware() {
	if os.Getenv("ENV") == "development" {
		// Set development only middleware here
		s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${method}\t${path}\t${status}\t${latency_human}\n",
		}))
	} else {
		// Set production only middleware here
		nrmiddleware, err := newRelicMiddleware()
		if err != nil {
			fmt.Println(err)
		} else {
			s.echo.Use(nrmiddleware)
		}
	}

	// Set shared middleware here
	s.echo.Use(middleware.Recover())
}
