package unit_test

import (
	"testing"

	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSerialsFavourites(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsFavourites)
	mockData := []*models.SerialsFavourites{
		{
			Sf_id:          1,
			Sf_idSerial:    1,
			Sf_idFavourite: 1,
		},
	}

	mockRepo.On("GetSerialsFavourites").Return(mockData, nil)

	serialsFavourites, err := mockRepo.GetSerialsFavourites()
	require.NoError(t, err)
	assert.NotNil(t, serialsFavourites)
	assert.Equal(t, mockData, serialsFavourites)
}

func TestGetSerialsFavouritesById(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsFavourites)
	mockData := &models.SerialsFavourites{
		Sf_id:          1,
		Sf_idSerial:    1,
		Sf_idFavourite: 1,
	}

	mockRepo.On("GetSerialsFavouritesById", 1).Return(mockData, nil)

	serialFavourite, err := mockRepo.GetSerialsFavouritesById(1)
	require.NoError(t, err)
	assert.NotNil(t, serialFavourite)
	assert.Equal(t, mockData, serialFavourite)
}

func TestCreateSerialsFavourites(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsFavourites)
	mockData := &models.SerialsFavourites{
		Sf_id:          1,
		Sf_idSerial:    1,
		Sf_idFavourite: 1,
	}

	mockRepo.On("CreateSerialsFavourites", mockData).Return(nil)

	err := mockRepo.CreateSerialsFavourites(mockData)
	require.NoError(t, err)
}

func TestUpdateSerialsFavourites(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsFavourites)
	mockData := &models.SerialsFavourites{
		Sf_id:          1,
		Sf_idSerial:    1,
		Sf_idFavourite: 1,
	}

	mockRepo.On("UpdateSerialsFavourites", mockData).Return(nil)

	err := mockRepo.UpdateSerialsFavourites(mockData)
	require.NoError(t, err)
}
