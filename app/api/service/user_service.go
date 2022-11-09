package service

import (
	"context"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo models.UserRepository
}

type UserConfig struct {
	UserRepo models.UserRepository
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*models.User, error) {
	u, err := s.UserRepo.FindByID(ctx, uid)
	return u, err
}

func CreateUserService(c *UserConfig) models.UserService {
	return &UserService{
		UserRepo: c.UserRepo,
	}
}
