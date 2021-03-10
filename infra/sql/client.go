package sql

import (
	"context"
	"fmt"
	"go-mysql-boilerplate/infra"
	"go-mysql-boilerplate/logger"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SqlDB holds necessery fields and
// mongo database session to connect
type SqlDB struct {
	database *gorm.DB
	name     string
	lgr      logger.Logger
}

// New returns a new instance of mongodb using session s
func New(ctx context.Context, driverName, dsn, name string, opts ...Option) (*SqlDB, error) {
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
	db := &SqlDB{
		database: database,
		name:     name,
	}
	for _, opt := range opts {
		opt.apply(db)
	}
	return db, nil
}

// Option is mongo db option
type Option interface {
	apply(*SqlDB)
}

// OptionFunc implements Option interface
type OptionFunc func(db *SqlDB)

func (f OptionFunc) apply(db *SqlDB) {
	f(db)
}

// SetLogger sets logger
func SetLogger(lgr logger.Logger) Option {
	return OptionFunc(func(db *SqlDB) {
		db.lgr = lgr
	})
}

func (d *SqlDB) println(args ...interface{}) {
	if d.lgr != nil {
		d.lgr.Println(args...)
	}
}

//func (d *SqlDB) Ping(ctx context.Context) error {
//	return d.database.Ping(ctx, readpref.Primary())
//}
//
//func (d *SqlDB) Close(ctx context.Context) error {
//	return d.Client.Disconnect(ctx)
//}

// EnsureIndices creates indices for collection col
func (d *SqlDB) EnsureIndices(ctx context.Context, inds []interface{}) error {
	log.Println("creating migration")
	var err error
	db := d.database
	for _, ind := range inds {
		err = db.AutoMigrate(ind)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *SqlDB) Insert(ctx context.Context, table string, v interface{}) error {
	panic("implement me")
}

func (d *SqlDB) InsertMany(ctx context.Context, table string, v []interface{}) error {
	panic("implement me")
}

func (d *SqlDB) FindMany(ctx context.Context, table string, filter interface{}, skip, limit int64, v interface{}, sort ...interface{}) error {
	panic("implement me")
}

func (d *SqlDB) FindOne(ctx context.Context, table string, filter infra.DbQuery, v interface{}, sort ...interface{}) error {
	panic("implement me")
}

func (d *SqlDB) UpdateMany(ctx context.Context, table string, filter infra.DbQuery, data interface{}) error {
	panic("implement me")
}

func (d *SqlDB) UpdateOne(ctx context.Context, table string, filter infra.DbQuery, data interface{}) error {
	panic("implement me")
}
