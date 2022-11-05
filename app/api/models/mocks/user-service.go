package mocks

import (
	"context"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(c context.Context, id uuid.UUID) (*models.User, error) {

	var rc0 *models.User
	r_call := m.Called(c, id)

	if r_call.Get(0) != nil {
		rc0 = r_call.Get(0).(*models.User)
	}

	var rc1 error
	if r_call.Get(1) != nil {
		rc1 = r_call.Get(1).(error)
	}

	return rc0, rc1
}
