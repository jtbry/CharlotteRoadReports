package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func newRelicMiddleware() (gin.HandlerFunc, error) {
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
	return nrgin.Middleware(app), nil
}

func (s *Server) setMiddleware() {
	if os.Getenv("ENV") == "development" {
		// Set development only middleware here
		s.router.Use(gin.Logger())
	} else {
		// Set production only middleware here
		nrmiddleware, err := newRelicMiddleware()
		if err != nil {
			fmt.Println(err)
		} else {
			s.router.Use(nrmiddleware)
		}
	}

	// Set shared middleware here
	s.router.Use(gin.Recovery())
	s.router.LoadHTMLGlob("web/templates/*.tmpl.html")
	s.router.Static("web/static", "static")
}
