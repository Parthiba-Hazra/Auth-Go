package models

import (
	"github.com/google/uuid"
)

// The handler need a method that return either User or an error
type UserService interface {
	Get(uid uuid.UUID) (*User, error)
}

// Service layer will expect this method
type UserRepository interface {
	FindByID(uid uuid.UUID) (*User, error)
}
