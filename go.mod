module github.com/jtbry/CharlotteRoadReports

// +heroku goVersion go1.15
go 1.15

require (
	cloud.google.com/go/storage v1.18.2 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.2.2
	github.com/newrelic/go-agent/v3 v3.11.0
	github.com/newrelic/go-agent/v3/integrations/nrecho-v4 v1.0.0
	github.com/sendgrid/rest v2.6.5+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.10.3+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	google.golang.org/api v0.58.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)
