package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// PostgreSQL holds sql config
type PostgreSQL struct {
	URL       string        `yaml:"url"`
	DBName    string        `yaml:"db_name"`
	DBTimeOut time.Duration `yaml:"time_out"`
}

var postgresqlOnce = sync.Once{}
var postgresqlConfig *PostgreSQL

// loadPostgreSQL loads config from path
func loadPostgreSQL(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	postgresqlConfig = &PostgreSQL{
		URL:       viper.GetString("sql.url"),
		DBName:    viper.GetString("sql.db_name"),
		DBTimeOut: viper.GetDuration("sql.time_out") * time.Second,
	}

	log.Println("sql config ", postgresqlConfig)
	return nil
}

// GetPostgreSQL returns sql config
func GetPostgreSQL(fileName string) *PostgreSQL {
	postgresqlOnce.Do(func() {
		loadPostgreSQL(fileName)
	})

	return postgresqlConfig
}
