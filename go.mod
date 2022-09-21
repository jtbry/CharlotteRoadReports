module github.com/jtbry/CharlotteRoadReports

// +heroku goVersion go1.15
go 1.15

require (
	cloud.google.com/go/iam v0.4.0 // indirect
	cloud.google.com/go/storage v1.22.1
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.2.2
	github.com/sendgrid/rest v2.6.5+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.10.3+incompatible
	github.com/spf13/viper v1.13.0
	google.golang.org/api v0.93.0
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)
