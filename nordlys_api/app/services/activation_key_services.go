package services

import "github.com/google/uuid"

// ActivationKeyServices service definition for activation key management.
type ActivationKeyServices interface {
	GenerateUniqueKey() string
}

// ActivationKeyServicesImpl Default implementation.
type ActivationKeyServicesImpl struct {}

// GenerateUniqueKey default implementation for unique key generation, using UUIDs.
func (service ActivationKeyServicesImpl) GenerateUniqueKey() string {
	id := uuid.New()
	return id.String()
}
