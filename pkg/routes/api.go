package routes

import (
	Controller "gitlab.com/jacob-ernst/mets/pkg/controllers/api"
	"gitlab.com/jacob-ernst/mets/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(api fiber.Router, db *database.Database) {
	registerActivities(api, db)
}

func registerActivities(api fiber.Router, db *database.Database) {
	activities := api.Group("/activities")

	activities.Get("/", Controller.GetAllActivities(db))
	activities.Get("/:id", Controller.GetActivity(db))
	activities.Post("/", Controller.AddActivity(db))
	activities.Put("/:id", Controller.EditActivity(db))
	activities.Delete("/:id", Controller.DeleteActivity(db))
}
