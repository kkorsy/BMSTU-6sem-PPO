package unit_test

import (
	"testing"

	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSerialsActors(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsActors)
	mockData := []*models.SerialsActors{
		{
			Sa_id:       1,
			Sa_idSerial: 1,
			Sa_idActor:  1,
		},
	}

	mockRepo.On("GetSerialsActors").Return(mockData, nil)

	serialsActors, err := mockRepo.GetSerialsActors()
	require.NoError(t, err)
	assert.NotNil(t, serialsActors)
	assert.Equal(t, mockData, serialsActors)
}

func TestGetSerialsByActorId(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsActors)
	mockData := []*models.SerialsActors{
		{
			Sa_id:       1,
			Sa_idSerial: 1,
			Sa_idActor:  1,
		},
	}

	mockRepo.On("GetSerialsByActorId", 1).Return(mockData, nil)

	serialsActors, err := mockRepo.GetSerialsByActorId(1)
	require.NoError(t, err)
	assert.NotNil(t, serialsActors)
	assert.Equal(t, mockData, serialsActors)
}

func TestGetActorsBySerialId(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsActors)
	mockData := []*models.SerialsActors{
		{
			Sa_id:       1,
			Sa_idSerial: 1,
			Sa_idActor:  1,
		},
	}

	mockRepo.On("GetActorsBySerialId", 1).Return(mockData, nil)

	serialsActors, err := mockRepo.GetActorsBySerialId(1)
	require.NoError(t, err)
	assert.NotNil(t, serialsActors)
	assert.Equal(t, mockData, serialsActors)
}

func TestGetSerialsActorsById(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsActors)
	mockData := &models.SerialsActors{
		Sa_id:       1,
		Sa_idSerial: 1,
		Sa_idActor:  1,
	}

	mockRepo.On("GetSerialsActorsById", 1).Return(mockData, nil)

	serialsActors, err := mockRepo.GetSerialsActorsById(1)
	require.NoError(t, err)
	assert.NotNil(t, serialsActors)
	assert.Equal(t, mockData, serialsActors)
}

func TestCreateSerialsActors(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsActors)
	mockData := &models.SerialsActors{
		Sa_id:       1,
		Sa_idSerial: 1,
		Sa_idActor:  1,
	}

	mockRepo.On("CreateSerialsActors", mockData).Return(nil)

	err := mockRepo.CreateSerialsActors(mockData)
	require.NoError(t, err)
}

func TestUpdateSerialsActors(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsActors)
	mockData := &models.SerialsActors{
		Sa_id:       1,
		Sa_idSerial: 1,
		Sa_idActor:  1,
	}

	mockRepo.On("UpdateSerialsActors", mockData).Return(nil)

	err := mockRepo.UpdateSerialsActors(mockData)
	require.NoError(t, err)
}
