package mocks

import (
	"app/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoSerials struct {
	mock.Mock
}

func (m *MockRepoSerials) GetSerials() ([]*models.Serial, error) {
	args := m.Called()
	return args.Get(0).([]*models.Serial), args.Error(1)
}

func (m *MockRepoSerials) GetSerialById(id int) (*models.Serial, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Serial), args.Error(1)
}

func (m *MockRepoSerials) GetSerialsByTitle(title string) ([]*models.Serial, error) {
	args := m.Called(title)
	return args.Get(0).([]*models.Serial), args.Error(1)
}

func (m *MockRepoSerials) CreateSerial(serial *models.Serial) error {
	args := m.Called(serial)
	return args.Error(0)
}

func (m *MockRepoSerials) UpdateSerial(serial *models.Serial) error {
	args := m.Called(serial)
	return args.Error(0)
}

func (m *MockRepoSerials) DeleteSerial(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepoSerials) CalculateDuration(serial *models.Serial) error {
	args := m.Called(serial)
	return args.Error(0)
}
