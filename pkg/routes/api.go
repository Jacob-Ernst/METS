package routes

import (
	Controller "gitlab.com/jacob-ernst/mets/pkg/controllers/api"
	"gitlab.com/jacob-ernst/mets/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(api fiber.Router, db *database.Database) {
	registerActivities(api, db)
	registerRoles(api, db)
	registerUsers(api, db)
}

func registerActivities(api fiber.Router, db *database.Database) {
	activities := api.Group("/activities")

	activities.Get("/", Controller.GetAllActivities(db))
	activities.Get("/:id", Controller.GetActivity(db))
	activities.Post("/", Controller.AddActivity(db))
	activities.Put("/:id", Controller.EditActivity(db))
	activities.Delete("/:id", Controller.DeleteActivity(db))
}

func registerRoles(api fiber.Router, db *database.Database) {
	roles := api.Group("/roles")

	roles.Get("/", Controller.GetAllRoles(db))
	roles.Get("/:id", Controller.GetRole(db))
	roles.Post("/", Controller.AddRole(db))
	roles.Put("/:id", Controller.EditRole(db))
	roles.Delete("/:id", Controller.DeleteRole(db))
}

func registerUsers(api fiber.Router, db *database.Database) {
	users := api.Group("/users")

	users.Get("/", Controller.GetAllUsers(db))
	users.Get("/:id", Controller.GetUser(db))
	users.Post("/", Controller.AddUser(db))
	users.Put("/:id", Controller.EditUser(db))
	users.Delete("/:id", Controller.DeleteUser(db))
}
