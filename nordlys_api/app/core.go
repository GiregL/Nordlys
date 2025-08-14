package app

import "gorm.io/gorm"

type Core struct {
	Database *gorm.DB
}
