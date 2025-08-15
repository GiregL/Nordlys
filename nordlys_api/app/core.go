package app

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Core struct {
	Database *gorm.DB
	Store    *session.Store
}
