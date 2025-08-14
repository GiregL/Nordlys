package resources

import (
	"nordlys_api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ActivationKey resource definition.
type ActivationKey struct {
	gorm.Model

	Key string
}

// ActivationKeyResource resource management definition.
type ActivationKeyResource struct {
	DB *gorm.DB
	ActivationKeyServices services.ActivationKeyServices
}

// GetAllActivationKey retrieves all the activation keys from the database.
func (resource *ActivationKeyResource) GetAllActivationKey() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// PostCreateActivationKey creates a new activation key into the database.
func (resource *ActivationKeyResource) PostCreateActivationKey() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}


