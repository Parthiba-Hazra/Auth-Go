package mocks

import (
	"context"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUsrRepo struct {
	mock.Mock
}

func (m *MockUsrRepo) FindByID(ctx context.Context, uid uuid.UUID) (*models.User, error) {

	res := m.Called(ctx, uid)

	var r0 *models.User
	if res.Get(0) != nil {
		r0 = res.Get(0).(*models.User)
	}

	var r1 error
	if res.Get(1) != nil {
		r1 = res.Get(1).(error)
	}

	return r0, r1
}
