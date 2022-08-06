package database

import (
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Cached   bool
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type Database struct {
	*gorm.DB
}

func New(config *DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error
	switch strings.ToLower(config.Driver) {
	case "postgresql", "postgres":
		dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite", "sqlite3":
		dsn := config.Database
		if config.Cached {
			dsn = dsn + "?cache=shared"
		}
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}
	return &Database{db}, err
}
