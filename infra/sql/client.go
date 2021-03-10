package sql

import (
	"context"
	"database/sql"
	"fmt"
	"go-mysql-boilerplate/logger"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB holds necessery fields and
// mongo Database session to connect
type DB struct {
	Database *gorm.DB
	SqlDB    *sql.DB
	Name     string
	lgr      logger.Logger
}

// New returns a new instance of mongodb using session s
func New(ctx context.Context, driverName, dsn, name string, opts ...Option) (*DB, error) {
	gormConfig := &gorm.Config{}

	var dialector gorm.Dialector
	if driverName == "sql" {
		dialector = mysql.New(mysql.Config{
			DriverName:        name,
			DSN:               dsn,
			Conn:              nil,
			DefaultStringSize: 255,
		})
	} else if driverName == "postgres" {
		dialector = postgres.New(postgres.Config{
			DriverName:           name,
			DSN:                  dsn,
			PreferSimpleProtocol: false,
			WithoutReturning:     false,
			Conn:                 nil,
		})
	} else {
		return nil, fmt.Errorf("driver not supported")
	}

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}

	db := &DB{
		Database: database,
		Name:     name,
		SqlDB:    sqlDB,
	}
	for _, opt := range opts {
		opt.apply(db)
	}
	return db, nil
}

// Option is mongo db option
type Option interface {
	apply(*DB)
}

// OptionFunc implements Option interface
type OptionFunc func(db *DB)

func (f OptionFunc) apply(db *DB) {
	f(db)
}

// SetLogger sets logger
func SetLogger(lgr logger.Logger) Option {
	return OptionFunc(func(db *DB) {
		db.lgr = lgr
	})
}

func (d *DB) println(args ...interface{}) {
	if d.lgr != nil {
		d.lgr.Println(args...)
	}
}

//func (d *DB) Ping(ctx context.Context) error {
//	return d.Database.Ping(ctx, readpref.Primary())
//}
//
//func (d *DB) Close(ctx context.Context) error {
//	return d.Client.Disconnect(ctx)
//}

// EnsureIndices creates indices for collection col
func (d *DB) EnsureIndices(ctx context.Context, inds []interface{}) error {
	log.Println("creating migration")
	var err error
	db := d.Database
	for _, ind := range inds {
		err = db.AutoMigrate(ind)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) Ping() error {
	log.Println("ping db")
	return d.SqlDB.Ping()
}
