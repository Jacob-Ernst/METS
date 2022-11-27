package main

import (
	"log"

	configuration "gitlab.com/jacob-ernst/mets/pkg/config"
	"gitlab.com/jacob-ernst/mets/pkg/database"
	"gitlab.com/jacob-ernst/mets/pkg/models"
	"gorm.io/gorm/clause"
)

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalln("failed to connect to database:", err.Error())
	}
	if db == nil {
		log.Fatalln("failed to connect to database: db variable is nil")
	}

	failedModel, err := migrateModels(db)
	if err != nil {
		log.Fatalln("failed to automigrate", failedModel, "model:", err.Error())
	}

	err = seedData(db)
	if err != nil {
		log.Fatalln("failed to seed data:", err.Error())
	}

	log.Println("ALL DONE!")
}

func migrateModels(db *database.Database) (string, error) {
	log.Println("automigrating models")
	err := db.AutoMigrate(&models.Activity{})
	if err != nil {
		return "activity", err
	}
	err = db.AutoMigrate(&models.Role{})
	if err != nil {
		return "role", err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return "user", err
	}

	return "", nil
}

func seedData(db *database.Database) error {
	log.Println("seeding activities")
	err := seedActivities(db)
	if err != nil {
		return err
	}

	log.Println("seeding roles")
	err = seedRoles(db)
	if err != nil {
		return err
	}

	return nil
}

func seedActivities(db *database.Database) error {
	tx := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&activityRecords)
	if tx.Error != nil {
		return tx.Error
	}

	log.Printf("inserted %d activities", tx.RowsAffected)
	return nil
}

func seedRoles(db *database.Database) error {
	tx := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roleRecords)
	if tx.Error != nil {
		return tx.Error
	}

	log.Printf("inserted %d roles", tx.RowsAffected)
	return nil
}

func setupDB() (*database.Database, error) {
	log.Println("setting up configuration")
	config := configuration.New()

	log.Println("connecting to database")
	// Initialize database
	db, err := database.New(&database.DatabaseConfig{
		Driver:   config.GetString("DB_DRIVER"),
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	return db, err
}
