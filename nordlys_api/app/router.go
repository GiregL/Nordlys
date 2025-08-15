package app

import (
	"nordlys_api/app/resources"
	"nordlys_api/app/services"

	"github.com/gofiber/fiber/v2"
)

// Routing initialisation and configuration.
func InitRouter(core *Core, app *fiber.App) {
	app.Static("/", "./public")

	group := app.Group("/api")

	// Activation key group
	activationKeyGroup := group.Group("/activation-key")
	InitActivationKeyRouter(core, app, activationKeyGroup)

	// User management and login group
	userGroup := group.Group("/user")
	InitUserRouter(core, app, userGroup)

}

func InitUserRouter(core *Core, app *fiber.App, group fiber.Router) {
	userResource := resources.UserResource{
		DB:    core.Database,
		Store: core.Store,
	}

	group.Get("/:id", userResource.GetUserById())
	group.Patch("/:id", userResource.PatchUpdateUserById())
	group.Delete("/:id", userResource.DeleteUserById())

	group.Post("/login", userResource.PostLoginUser())
	group.Post("/register", userResource.PostRegisterNewUser())
}

// InitActivationKeyRouter initializes the router part for the activation key resource management.
func InitActivationKeyRouter(core *Core, app *fiber.App, group fiber.Router) {
	service := services.ActivationKeyServicesImpl{}

	activationKeyResource := resources.ActivationKeyResource{
		DB:                    core.Database,
		ActivationKeyServices: service,
	}

	group.Get("/", activationKeyResource.GetAllActivationKey())
	group.Post("/", activationKeyResource.PostCreateActivationKey())
	group.Delete("/:id", activationKeyResource.DeleteActivationKey())
}

func defaultHandler(c *fiber.Ctx) error {
	return c.SendString("Not yet implemented")
}
