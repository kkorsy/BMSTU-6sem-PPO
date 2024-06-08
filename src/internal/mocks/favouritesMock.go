package mocks

import (
	"app/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoFavourites struct {
	mock.Mock
}

func (m *MockRepoFavourites) GetFavourites() ([]*models.Favourites, error) {
	args := m.Called()
	return args.Get(0).([]*models.Favourites), args.Error(1)
}

func (m *MockRepoFavourites) GetFavouriteById(id int) (*models.Favourites, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Favourites), args.Error(1)
}

func (m *MockRepoFavourites) CreateFavourite(favourite *models.Favourites) (int, error) {
	args := m.Called(favourite)
	return args.Int(0), args.Error(1)
}

func (m *MockRepoFavourites) UpdateFavourite(favourite *models.Favourites) error {
	args := m.Called(favourite)
	return args.Error(0)
}

func (m *MockRepoFavourites) DeleteFavourite(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
