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


}

// InitActivationKeyRouter initializes the router part for the activation key resource management.
func InitActivationKeyRouter(core *Core, app *fiber.App, group fiber.Router) {
	service := services.ActivationKeyServicesImpl{}

	activationKeyResource := resources.ActivationKeyResource{
		DB: core.Database,
		ActivationKeyServices: service,
	}

	group.Get("/", activationKeyResource.GetAllActivationKey())
	group.Post("/", activationKeyResource.PostCreateActivationKey())
}

func defaultHandler(c *fiber.Ctx) error {
	return c.SendString("Not yet implemented")
}
