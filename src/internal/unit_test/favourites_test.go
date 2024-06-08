package unit_test

import (
	"testing"

	"app/internal/controllers"
	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFavourites(t *testing.T) {
	mockRepo := new(mocks.MockRepoFavourites)
	mockRepo.On("GetFavourites").Return([]*models.Favourites{{F_id: 1}}, nil)

	ctrl := controllers.NewFavouritesCtrl(mockRepo)
	favourites, err := ctrl.GetFavourites()

	require.NoError(t, err)
	assert.Equal(t, 1, favourites[0].F_id)
	mockRepo.AssertCalled(t, "GetFavourites")
}

func TestGetFavouriteById(t *testing.T) {
	mockRepo := new(mocks.MockRepoFavourites)
	mockRepo.On("GetFavouriteById", 1).Return(&models.Favourites{F_id: 1}, nil)

	ctrl := controllers.NewFavouritesCtrl(mockRepo)
	favourite, err := ctrl.GetFavouriteById(1)

	require.NoError(t, err)
	assert.Equal(t, 1, favourite.F_id)
	mockRepo.AssertCalled(t, "GetFavouriteById", 1)
}

func TestUpdateFavourite(t *testing.T) {
	mockRepo := new(mocks.MockRepoFavourites)
	mockRepo.On("UpdateFavourite", &models.Favourites{F_id: 1}).Return(nil)

	ctrl := controllers.NewFavouritesCtrl(mockRepo)
	err := ctrl.UpdateFavourite(&models.Favourites{F_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateFavourite", &models.Favourites{F_id: 1})
}
