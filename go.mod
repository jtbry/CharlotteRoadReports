module github.com/jtbry/CharlotteRoadReports

// +heroku goVersion go1.15
go 1.15

require (
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.2.2
	github.com/newrelic/go-agent/v3 v3.11.0
	github.com/newrelic/go-agent/v3/integrations/nrecho-v4 v1.0.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)
