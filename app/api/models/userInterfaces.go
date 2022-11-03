package models

import (
	"context"

	"github.com/google/uuid"
)

// The handler need a method that return either User or an error
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

// Service layer will expect this method
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}
