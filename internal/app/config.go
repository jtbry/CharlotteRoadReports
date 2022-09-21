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
	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, err
	}

	// Unmarshal config
	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		return AppConfig{}, err
	}

	// Validate config
	if config.DatabaseURL == "" {
		return AppConfig{}, errors.New("DATABASE_URL is required")
	}

	if config.ServeClient {
		if _, err := os.Stat("./web/build"); os.IsNotExist(err) {
			return AppConfig{}, errors.New("web/build does not exist")
		}
	}

	return config, nil
}
