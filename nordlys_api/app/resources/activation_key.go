package resources

import (
	"context"
	"nordlys_api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ActivationKey resource definition.
type ActivationKey struct {
	gorm.Model

	Key string `json:"key"`
	Enabled bool `json:"enabled"`
}

// ActivationKeyResource resource management definition.
type ActivationKeyResource struct {
	DB *gorm.DB
	ActivationKeyServices services.ActivationKeyServices
}

// GetAllActivationKey retrieves all the activation keys from the database.
func (resource *ActivationKeyResource) GetAllActivationKey() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var keys []ActivationKey
		result := resource.DB.Find(&keys)

		if result.Error != nil {
			return ctx.SendStatus(500)
		} else {
			return ctx.JSON(keys)
		}
	}
}

// PostCreateActivationKey creates a new activation key from a generated UUID.
func (resource *ActivationKeyResource) PostCreateActivationKey() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		uuid := resource.ActivationKeyServices.GenerateUniqueKey()

		activationKey := ActivationKey{
			Key: uuid,
			Enabled: true,
		}

		dbCtx := context.Background()
		err := gorm.G[ActivationKey](resource.DB).Create(dbCtx, &activationKey)

		if err != nil {
			return c.SendStatus(500)
		} else {
			return c.SendStatus(201)
		}
	}
}

// DeleteActivationKey deletes an activation key from the system.
func (resource *ActivationKeyResource) DeleteActivationKey() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		keyId := c.Params("id")

		dbCtx := context.Background()

		affected, err := gorm.G[ActivationKey](resource.DB).Where("ID = ?", keyId).Delete(dbCtx)
		if affected == 0 {
			return c.SendStatus(404)
		} else if err != nil {
			return c.SendStatus(500)
		} else {
			return c.SendStatus(204)
		}
	}
}

// type ActivationKeyCreate struct {
// 	Key string `json:"key" xml:"key" form:"key"`
// }
//
// // PostCreateActivationKey creates a new activation key into the database.
// func (resource *ActivationKeyResource) PostCreateActivationKey() func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		form := new(ActivationKeyCreate)
// 		if err := c.BodyParser(form); err != nil {
// 			return err
// 		}
//
// 		activationKey := ActivationKey{
// 			Key: form.Key,
// 		}
//
// 		dbCtx := context.Background()
// 		err := gorm.G[ActivationKey](resource.DB).Create(dbCtx, &activationKey)
//
// 		if err != nil {
// 			return c.SendStatus(500)
// 		} else {
// 			return c.SendStatus(201)
// 		}
// 	}
// }


