package app

import "github.com/gofiber/fiber/v2"

// Routing initialisation and configuration.
func InitRouter(app *fiber.App) {

	group := app.Group("/api")

	// Authentication
	authGroup := group.Group("/auth")
	authGroup.Post("/register", defaultHandler)
	authGroup.Post("/login", defaultHandler)
	authGroup.Post("/logout", defaultHandler)

	// User
	userGroup := group.Group("/user")
	userGroup.Get("/profile/:id", defaultHandler)
	userGroup.Put("/profile/:id", defaultHandler)
	userGroup.Delete("/profile/:id", defaultHandler)

	// Medical calendar
	medicalCalendarGroup := group.Group("/medical-calendar")

	// Mood tracking
	moodTrackingGroup := group.Group("/mood-tracking")

	// Flashcards, notes, links
	languageGroup := group.Group("/languages")

	// Budget
	budgetGroup := group.Group("/budget")

	// General

}

func defaultHandler(c *fiber.Ctx) error {
	return c.SendString("Not yet implemented")
}
