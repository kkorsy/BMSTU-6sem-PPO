package mocks

import (
	"app/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoSeasons struct {
	mock.Mock
}

func (m *MockRepoSeasons) GetSeasons() ([]*models.Seasons, error) {
	args := m.Called()
	return args.Get(0).([]*models.Seasons), args.Error(1)
}

func (m *MockRepoSeasons) GetSeasonById(id int) (*models.Seasons, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Seasons), args.Error(1)
}

func (m *MockRepoSeasons) GetSeasonsBySerialId(id int) ([]*models.Seasons, error) {
	args := m.Called(id)
	return args.Get(0).([]*models.Seasons), args.Error(1)
}

func (m *MockRepoSeasons) CreateSeason(season *models.Seasons) error {
	args := m.Called(season)
	return args.Error(0)
}

func (m *MockRepoSeasons) UpdateSeason(season *models.Seasons) error {
	args := m.Called(season)
	return args.Error(0)
}

func (m *MockRepoSeasons) DeleteSeason(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
