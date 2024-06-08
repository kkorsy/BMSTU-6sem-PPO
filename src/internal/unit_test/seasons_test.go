package unit_test

import (
	"testing"

	"app/internal/controllers"
	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSeasons(t *testing.T) {
	mockRepo := new(mocks.MockRepoSeasons)
	mockRepo.On("GetSeasons").Return([]*models.Seasons{{Ss_id: 1}}, nil)

	ctrl := controllers.NewSeasonsCtrl(mockRepo)
	seasons, err := ctrl.GetSeasons()

	require.NoError(t, err)
	assert.Equal(t, 1, seasons[0].Ss_id)
	mockRepo.AssertCalled(t, "GetSeasons")
}

func TestGetSeasonById(t *testing.T) {
	mockRepo := new(mocks.MockRepoSeasons)
	mockRepo.On("GetSeasonById", 1).Return(&models.Seasons{Ss_id: 1}, nil)

	ctrl := controllers.NewSeasonsCtrl(mockRepo)
	season, err := ctrl.GetSeasonById(1)

	require.NoError(t, err)
	assert.Equal(t, 1, season.Ss_id)
	mockRepo.AssertCalled(t, "GetSeasonById", 1)
}

func TestCreateSeason(t *testing.T) {
	mockRepo := new(mocks.MockRepoSeasons)
	mockRepo.On("CreateSeason", &models.Seasons{Ss_id: 1}).Return(nil)

	ctrl := controllers.NewSeasonsCtrl(mockRepo)
	err := ctrl.CreateSeason(&models.Seasons{Ss_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateSeason", &models.Seasons{Ss_id: 1})
}

func TestUpdateSeason(t *testing.T) {
	mockRepo := new(mocks.MockRepoSeasons)
	mockRepo.On("UpdateSeason", &models.Seasons{Ss_id: 1}).Return(nil)

	ctrl := controllers.NewSeasonsCtrl(mockRepo)
	err := ctrl.UpdateSeason(&models.Seasons{Ss_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateSeason", &models.Seasons{Ss_id: 1})
}

func TestDeleteSeason(t *testing.T) {
	mockRepo := new(mocks.MockRepoSeasons)
	mockRepo.On("DeleteSeason", 1).Return(nil)

	ctrl := controllers.NewSeasonsCtrl(mockRepo)
	err := ctrl.DeleteSeason(1)

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteSeason", 1)
}
