package mocks

import (
	"app/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockRepoEpisodes struct {
	mock.Mock
}

func (m *MockRepoEpisodes) GetEpisodes() ([]*models.Episodes, error) {
	args := m.Called()
	return args.Get(0).([]*models.Episodes), args.Error(1)
}

func (m *MockRepoEpisodes) GetEpisodeById(id int) (*models.Episodes, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Episodes), args.Error(1)
}

func (m *MockRepoEpisodes) CreateEpisode(episode *models.Episodes) error {
	args := m.Called(episode)
	return args.Error(0)
}

func (m *MockRepoEpisodes) UpdateEpisode(episode *models.Episodes) error {
	args := m.Called(episode)
	return args.Error(0)
}

func (m *MockRepoEpisodes) DeleteEpisode(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
