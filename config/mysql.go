package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// MySQL holds sql config
type MySQL struct {
	URL       string        `yaml:"url"`
	DBName    string        `yaml:"db_name"`
	DBTimeOut time.Duration `yaml:"time_out"`
}

var mysqlOnce = sync.Once{}
var mysqlConfig *MySQL

// loadMySQL loads config from path
func loadMySQL(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	mysqlConfig = &MySQL{
		URL:       viper.GetString("sql.url"),
		DBName:    viper.GetString("sql.db_name"),
		DBTimeOut: viper.GetDuration("sql.time_out") * time.Second,
	}

	log.Println("sql config ", mysqlConfig)
	return nil
}

// GetMySQL returns sql config
func GetMySQL(fileName string) *MySQL {
	mysqlOnce.Do(func() {
		loadMySQL(fileName)
	})

	return mysqlConfig
}
