package unit_test

import (
	"testing"

	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProducers(t *testing.T) {
	mockRepo := new(mocks.MockRepoProducers)
	mockData := []*models.Producers{
		{
			P_id:   1,
			P_name: "Test",
		},
	}

	mockRepo.On("GetProducers").Return(mockData, nil)

	producers, err := mockRepo.GetProducers()
	require.NoError(t, err)
	assert.NotNil(t, producers)
	assert.Equal(t, mockData, producers)
}

func TestGetProducerById(t *testing.T) {
	mockRepo := new(mocks.MockRepoProducers)
	mockData := &models.Producers{
		P_id:   1,
		P_name: "Test",
	}

	mockRepo.On("GetProducerById", 1).Return(mockData, nil)

	producer, err := mockRepo.GetProducerById(1)
	require.NoError(t, err)
	assert.NotNil(t, producer)
	assert.Equal(t, mockData, producer)
}

func TestCreateProducer(t *testing.T) {
	mockRepo := new(mocks.MockRepoProducers)
	mockData := &models.Producers{
		P_id:   1,
		P_name: "Test",
	}

	mockRepo.On("CreateProducer", mockData).Return(nil)

	err := mockRepo.CreateProducer(mockData)
	require.NoError(t, err)
}

func TestUpdateProducer(t *testing.T) {
	mockRepo := new(mocks.MockRepoProducers)
	mockData := &models.Producers{
		P_id:   1,
		P_name: "Test",
	}

	mockRepo.On("UpdateProducer", mockData).Return(nil)

	err := mockRepo.UpdateProducer(mockData)
	require.NoError(t, err)
}

func TestDeleteProducer(t *testing.T) {
	mockRepo := new(mocks.MockRepoProducers)

	mockRepo.On("DeleteProducer", 1).Return(nil)

	err := mockRepo.DeleteProducer(1)
	require.NoError(t, err)
}
