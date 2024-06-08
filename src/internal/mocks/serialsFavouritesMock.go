package mocks

import (
	"app/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoSerialsFavourites struct {
	mock.Mock
}

func (m *MockRepoSerialsFavourites) GetSerialsFavourites() ([]*models.SerialsFavourites, error) {
	args := m.Called()
	return args.Get(0).([]*models.SerialsFavourites), args.Error(1)
}

func (m *MockRepoSerialsFavourites) GetSerialsFavouritesById(id int) (*models.SerialsFavourites, error) {
	args := m.Called(id)
	return args.Get(0).(*models.SerialsFavourites), args.Error(1)
}

func (m *MockRepoSerialsFavourites) CreateSerialsFavourites(serialFavourite *models.SerialsFavourites) error {
	args := m.Called(serialFavourite)
	return args.Error(0)
}

func (m *MockRepoSerialsFavourites) UpdateSerialsFavourites(serialFavourite *models.SerialsFavourites) error {
	args := m.Called(serialFavourite)
	return args.Error(0)
}

func (m *MockRepoSerialsFavourites) DeleteSerialsFavourites(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
