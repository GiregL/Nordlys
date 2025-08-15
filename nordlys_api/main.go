package main

import (
	"log"
	nordlys_api "nordlys_api/app"
	"nordlys_api/app/resources"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Database setup
	dsn := "host=localhost user=postgres password=postgres dbname=nordlys_dev port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("An error occured while connecting to database: ", err.Error())
	}

	// Database model migrations.
	err = db.AutoMigrate(resources.User{}, resources.ActivationKey{})
	if err != nil {
		log.Fatal("An error occured while migrating database: ", err.Error())
	}

	core := nordlys_api.Core{
		Database: db,
		Store:    session.New(),
	}

	nordlys_api.InitRouter(&core, app)

	log.Fatal(app.Listen(":3000"))
}
