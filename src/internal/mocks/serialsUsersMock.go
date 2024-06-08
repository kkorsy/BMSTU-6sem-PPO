package mocks

import (
	"app/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockRepoSerialsUsers struct {
	mock.Mock
}

func (m *MockRepoSerialsUsers) GetSerialsUsers() ([]*models.SerialsUsers, error) {
	args := m.Called()
	return args.Get(0).([]*models.SerialsUsers), args.Error(1)
}

func (m *MockRepoSerialsUsers) GetSerialsUsersById(id int) (*models.SerialsUsers, error) {
	args := m.Called(id)
	return args.Get(0).(*models.SerialsUsers), args.Error(1)
}

func (m *MockRepoSerialsUsers) GetSerialsByUserId(id int) ([]*models.SerialsUsers, error) {
	args := m.Called(id)
	return args.Get(0).([]*models.SerialsUsers), args.Error(1)
}

func (m *MockRepoSerialsUsers) GetUsersBySerialId(id int) ([]*models.SerialsUsers, error) {
	args := m.Called(id)
	return args.Get(0).([]*models.SerialsUsers), args.Error(1)
}

func (m *MockRepoSerialsUsers) CreateSerialsUsers(serialUser *models.SerialsUsers) error {
	args := m.Called(serialUser)
	return args.Error(0)
}

func (m *MockRepoSerialsUsers) UpdateSerialsUsers(serialUser *models.SerialsUsers) error {
	args := m.Called(serialUser)
	return args.Error(0)
}
