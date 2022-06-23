package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB(dbName string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&Activity{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
