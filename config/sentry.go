package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

// Sentry holds sentry config
type Sentry struct {
	URL string `yaml:"sentry.dsn"`
}

var sentryOnce = sync.Once{}
var sentryConfig *Sentry

// loadSentry loads config from path
func loadSentry(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	sentryConfig = &Sentry{
		URL: viper.GetString("sentry.dsn"),
	}

	log.Println("sentry config ", tableConfig)
	return nil
}

// GetSentry returns redis config
func GetSentry(fileName string) *Sentry {
	sentryOnce.Do(func() {
		loadSentry(fileName)
	})

	return sentryConfig
}
