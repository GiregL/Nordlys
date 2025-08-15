package resources

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

// User resource definition.
type User struct {
	gorm.Model

	Username string `json:"Username"`
	Password string

	UserActivationKeyID int `json:"UserActivationKeyID"`
	UserActivationKey   ActivationKey
}

type UserResource struct {
	DB    *gorm.DB
	Store *session.Store
}

type UserLoginForm struct {
	Username string `json:"Username" form:"Username" binding:"required"`
	Password string `json:"Password" form:"Password" binding:"required"`
}

// PostLoginUser logins a user.
// The authentication system relies on sessions.
func (u *UserResource) PostLoginUser() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		form := new(UserLoginForm)
		if err := c.BodyParser(form); err != nil {
			return err
		}

		dbCtx := context.Background()
		user, err := gorm.G[User](u.DB).Where("Username = ? AND Password = ?", form.Username, form.Password).First(dbCtx)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		sess, err := u.Store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		sess.Set("user_id", user.ID)

		return c.JSON(user)
	}
}

// GetUserById retrieves the user by its unique ID.
func (u *UserResource) GetUserById() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(500)
	}
}

// PostRegisterNewUser corresponds to the user registration in the system.
func (u *UserResource) PostRegisterNewUser() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(500)
	}
}

// PatchUpdateUserById updates the given user information by its ID.
func (u *UserResource) PatchUpdateUserById() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(500)
	}
}

// DeleteUserById deletes a user by its ID (physical deletion)
func (u *UserResource) DeleteUserById() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(500)
	}
}
