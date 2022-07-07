package api

import (
	"gitlab.com/jacob-ernst/mets/pkg/database"
	"gitlab.com/jacob-ernst/mets/pkg/models"

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

// Add a single activity to the database
func AddActivity(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Activity := new(models.Activity)
		if err := ctx.BodyParser(Activity); err != nil {
			panic("An error occurred when parsing the new activity: " + err.Error())
		}
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

// Edit a single activity
func EditActivity(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditActivity := new(models.Activity)
		Activity := new(models.Activity)
		if err := ctx.BodyParser(EditActivity); err != nil {
			panic("An error occurred when parsing the edited activity: " + err.Error())
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
		Activity.Name = EditActivity.Name
		Activity.Description = EditActivity.Description
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
