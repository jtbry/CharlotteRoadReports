package app

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Env               string `mapstructure:"ENV"`
	Port              string `mapstructure:"PORT"`
	DatabaseURL       string `mapstructure:"DATABASE_URL"`
	ServeClient       bool   `mapstructure:"SERVE_CLIENT"`
	ScheduledScraping bool   `mapstructure:"SCHEDULED_SCRAPING"`
}

func (c *AppConfig) validate() error {
	// Validate that a DSN is given
	if c.DatabaseURL == "" {
		return errors.New("DATABASE_URL is required")
	}

	// Validate client build files exist if serving client
	if c.ServeClient {
		if _, err := os.Stat("./web/build"); os.IsNotExist(err) {
			return errors.New("web/build does not exist")
		}
	}

	return nil
}

func LoadConfig() (AppConfig, error) {
	// Set defaults
	viper.SetDefault("ENV", "development")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("SERVE_CLIENT", false)
	viper.SetDefault("SCHEDULED_SCRAPING", false)

	// Set config file
	viper.AddConfigPath(".")
	if os.Getenv("ENV") == "production" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("dev")
	}
	viper.SetConfigType("env")

	// Read config
	viper.AutomaticEnv()
	viper.ReadInConfig()

	// Unmarshal config
	var config AppConfig
	viper.Unmarshal(&config)

	// Validate config
	err := config.validate()
	if err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
