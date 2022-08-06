package api

import (
	"gitlab.com/jacob-ernst/mets/pkg/database"
	"gitlab.com/jacob-ernst/mets/pkg/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Return all activities as JSON
func GetAllActivities(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Activity []models.Activity
		if response := db.Find(&Activity); response.Error != nil {
			panic("Error occurred while retrieving activities from the database: " + response.Error.Error())
		}
		err := ctx.JSON(Activity)
		if err != nil {
			panic("Error occurred when returning JSON of activities: " + err.Error())
		}
		return err
	}
}

// Return a single activity as JSON
func GetActivity(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Activity := new(models.Activity)
		id := ctx.Params("id")
		if response := db.Find(&Activity, id); response.Error != nil {
			panic("An error occurred when retrieving the activity: " + response.Error.Error())
		}
		if Activity.ID == 0 {
			// Send status not found
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				panic("Cannot return status not found: " + err.Error())
			}
			// Set ID
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				panic("Error occurred when returning JSON of a activity: " + err.Error())
			}
			return err
		}
		err := ctx.JSON(Activity)
		if err != nil {
			panic("Error occurred when returning JSON of a activity: " + err.Error())
		}
		return err
	}
}

type AddParams struct {
	Effort      float64 `validate:"required,gte=0.5"`
	Name        string  `validate:"required"`
	Description string
}

// Add a single activity to the database
func AddActivity(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		validator := validator.New()
		params := new(AddParams)
		Activity := new(models.Activity)
		if err := ctx.BodyParser(params); err != nil {
			panic("An error occurred when parsing the new activity: " + err.Error())
		}
		if err := validator.Struct(params); err != nil {
			panic("An error occurred when validating the new activity: " + err.Error())
		}

		Activity.Name = params.Name
		Activity.Description = params.Description
		Activity.Effort = params.Effort

		if response := db.Create(&Activity); response.Error != nil {
			panic("An error occurred when storing the new activity: " + response.Error.Error())
		}
		err := ctx.JSON(Activity)
		if err != nil {
			panic("Error occurred when returning JSON of a activity: " + err.Error())
		}
		return err
	}
}

type EditParams struct {
	ID          uint    `validate:"required,gte=1"`
	Effort      float64 `validate:"required,gte=0.5"`
	Name        string  `validate:"required"`
	Description string
}

// Edit a single activity
func EditActivity(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		validator := validator.New()
		id := ctx.Params("id")
		params := new(EditParams)
		Activity := new(models.Activity)
		if err := ctx.BodyParser(params); err != nil {
			panic("An error occurred when parsing the edited activity: " + err.Error())
		}
		if err := validator.Struct(params); err != nil {
			panic("An error occurred when validating the new activity: " + err.Error())
		}
		if response := db.Find(&Activity, id); response.Error != nil {
			panic("An error occurred when retrieving the existing activity: " + response.Error.Error())
		}
		// Activity does not exist
		if Activity.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				panic("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				panic("Error occurred when returning JSON of a activity: " + err.Error())
			}
			return err
		}
		Activity.Name = params.Name
		Activity.Description = params.Description
		Activity.Effort = params.Effort
		db.Save(&Activity)

		err := ctx.JSON(Activity)
		if err != nil {
			panic("Error occurred when returning JSON of a activity: " + err.Error())
		}
		return err
	}
}

// Delete a single activity
func DeleteActivity(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var Activity models.Activity
		db.Find(&Activity, id)
		if response := db.Find(&Activity); response.Error != nil {
			panic("An error occurred when finding the activity to be deleted: " + response.Error.Error())
		}
		db.Delete(&Activity)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			panic("Error occurred when returning JSON of a activity: " + err.Error())
		}
		return err
	}
}
