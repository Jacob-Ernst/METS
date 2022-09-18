package main

import (
	"log"

	configuration "gitlab.com/jacob-ernst/mets/pkg/config"
	"gitlab.com/jacob-ernst/mets/pkg/database"
	"gitlab.com/jacob-ernst/mets/pkg/models"
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

	return nil
}

func seedActivities(db *database.Database) error {
	var count int
	tx := db.Exec("INSERT INTO activities (name, effort) VALUES ('power mower', 4.5) ON CONFLICT (name) DO NOTHING")
	if tx.Error != nil {
		return tx.Error
	}
	count++

	tx = db.Exec("INSERT INTO activities (name, description, effort) VALUES ('running, 4 mph', 'Running 15 min/mile', 6) ON CONFLICT (name) DO NOTHING")
	if tx.Error != nil {
		return tx.Error
	}
	count++

	tx = db.Exec("INSERT INTO activities (name, description, effort) VALUES ('sitting tasks, light effort', 'Examples are office work, chemistry lab work, computer work, light assembly repair, watch repair, reading, desk work', 1.5) ON CONFLICT (name) DO NOTHING")
	if tx.Error != nil {
		return tx.Error
	}
	count++

	log.Printf("inserted or updated %d activities", count)
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
