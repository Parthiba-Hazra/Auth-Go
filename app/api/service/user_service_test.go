package service

import (
	"context"
	"testing"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/Parthiba-Hazra/auth-go/models/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TryGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserRes := &models.User{
			UID:   uid,
			Email: "rono@gmail.com",
			Name:  "Cris Cris",
		}

		mockUserRepo := new(mocks.MockUsrRepo)
		usrService := CreateUserService(&UserConfig{
			UserRepo: mockUserRepo,
		})
		mockUserRepo.On("FindByID", mock.Anything, uid).Return(mockUserRes, nil)

		ctx := context.TODO()
		u, err := usrService.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, u, mockUserRepo)
		mockUserRepo.AssertExpectations(t)
	})
}
