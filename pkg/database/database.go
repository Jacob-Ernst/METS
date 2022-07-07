package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DSN string
}

type Database struct {
	*gorm.DB
}

func New(config *DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	return &Database{db}, err
}
