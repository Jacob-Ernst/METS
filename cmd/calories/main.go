package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gitlab.com/jacob-ernst/mets/pkg/calories"
	"gitlab.com/jacob-ernst/mets/pkg/conversions"
	"gitlab.com/jacob-ernst/mets/pkg/models"
)

var db *gorm.DB

func main() {
	var activity string
	var kg, lb, MET, minutes float64
	flag.StringVar(&activity, "activity", "", "the activity you engaged in")
	flag.Float64Var(&kg, "kg", -1, "your weight in KG")
	flag.Float64Var(&lb, "lb", -1, "your weight in pounds")
	flag.Float64Var(&MET, "met", -1, "MET for the task")
	flag.Float64Var(&minutes, "time", 20.00, "time you spent in minutes")
	flag.Parse()

	err := openDB("tmp/dev.db")
	if err != nil {
		log.Fatalln(err)
	}

	effort, err := getMET(activity, MET)
	if err != nil {
		log.Fatalln(err)
	}

	weight, err := convertWeight(kg, lb)
	if err != nil {
		log.Fatalln(err)
	}

	total := calories.TotalBurn(weight, effort, minutes)

	fmt.Printf("You burned %.2f Calories\n", total)
}

// Returns weight in kg when able, otherwise converts lbs to kg.
func convertWeight(kg, lb float64) (float64, error) {
	if kg >= 1 {
		return kg, nil
	}

	if lb < 1 {
		return -1, errors.New("Must provide either -kg or -lb.")
	}

	return conversions.PoundsToMetric(lb), nil
}

func getMET(activity string, MET float64) (float64, error) {
	if MET >= 1 {
		return MET, nil
	}

	if activity == "" {
		return -1, errors.New("Must provide either -activity or -met.")
	}

	var value models.Activity
	result := db.First(&value, "name = ?", activity)
	if result.Error != nil {
		return -1, result.Error
	}

	return value.Effort, nil
}

func openDB(dbName string) (err error) {
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Activity{})
	if err != nil {
		return err
	}

	// Create
	db.Create(&models.Activity{Name: "power mower", Effort: 4.5})
	return nil
}
