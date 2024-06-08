package unit_test

import (
	"testing"

	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEpisodes(t *testing.T) {
	mockRepo := new(mocks.MockRepoEpisodes)
	mockData := []*models.Episodes{
		{
			E_id:   1,
			E_name: "Test Title",
		},
	}

	mockRepo.On("GetEpisodes").Return(mockData, nil)

	episodes, err := mockRepo.GetEpisodes()
	require.NoError(t, err)
	assert.NotNil(t, episodes)
	assert.Equal(t, mockData, episodes)
}

func TestGetEpisodeById(t *testing.T) {
	mockRepo := new(mocks.MockRepoEpisodes)
	mockData := &models.Episodes{
		E_id:   1,
		E_name: "Test Title",
	}

	mockRepo.On("GetEpisodeById", 1).Return(mockData, nil)

	episode, err := mockRepo.GetEpisodeById(1)
	require.NoError(t, err)
	assert.NotNil(t, episode)
	assert.Equal(t, mockData, episode)
}

func TestCreateEpisode(t *testing.T) {
	mockRepo := new(mocks.MockRepoEpisodes)
	mockData := &models.Episodes{
		E_id:   1,
		E_name: "Test Title",
	}

	mockRepo.On("CreateEpisode", mockData).Return(nil)

	err := mockRepo.CreateEpisode(mockData)
	require.NoError(t, err)
}

func TestUpdateEpisode(t *testing.T) {
	mockRepo := new(mocks.MockRepoEpisodes)
	mockData := &models.Episodes{
		E_id:   1,
		E_name: "Test Title",
	}

	mockRepo.On("UpdateEpisode", mockData).Return(nil)

	err := mockRepo.UpdateEpisode(mockData)
	require.NoError(t, err)
}

func TestDeleteEpisode(t *testing.T) {
	mockRepo := new(mocks.MockRepoEpisodes)

	mockRepo.On("DeleteEpisode", 1).Return(nil)

	err := mockRepo.DeleteEpisode(1)
	require.NoError(t, err)
}
