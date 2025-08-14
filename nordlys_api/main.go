package main

import (
	nordlys_api "nordlys_api/app"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	core := nordlys_api.Core{}
	nordlys_api.InitRouter(&core, app)

	log.Fatal(app.Listen(":3000"))
}
