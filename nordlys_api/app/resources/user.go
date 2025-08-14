package resources

import "gorm.io/gorm"

// User resource definition.
type User struct {
	gorm.Model

	Username string
	Password string

	UserActivationKeyID int
	UserActivationKey ActivationKey
}
