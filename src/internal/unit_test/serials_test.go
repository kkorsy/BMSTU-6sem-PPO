package unit_test

import (
	"testing"

	"app/internal/controllers"
	"app/internal/mocks"
	"app/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSerials(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerials)
	mockRepo.On("GetSerials").Return([]*models.Serial{{S_id: 1}}, nil)

	ctrl := controllers.NewSerialsCtrl(mockRepo)
	serials, err := ctrl.GetSerials()

	require.NoError(t, err)
	assert.Equal(t, 1, serials[0].S_id)
	mockRepo.AssertCalled(t, "GetSerials")
}

func TestGetSerialById(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerials)
	mockRepo.On("GetSerialById", 1).Return(&models.Serial{S_id: 1}, nil)

	ctrl := controllers.NewSerialsCtrl(mockRepo)
	serial, err := ctrl.GetSerialById(1)

	require.NoError(t, err)
	assert.Equal(t, 1, serial.S_id)
	mockRepo.AssertCalled(t, "GetSerialById", 1)
}

func TestCreateSerial(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerials)
	mockRepo.On("CreateSerial", &models.Serial{S_id: 1}).Return(nil)
	mockRepo.On("CalculateDuration", &models.Serial{S_id: 1}).Return(nil)

	ctrl := controllers.NewSerialsCtrl(mockRepo)
	err := ctrl.CreateSerial(&models.Serial{S_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateSerial", &models.Serial{S_id: 1})
	mockRepo.AssertCalled(t, "CalculateDuration", &models.Serial{S_id: 1})
}

func TestUpdateSerial(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerials)
	mockRepo.On("UpdateSerial", &models.Serial{S_id: 1}).Return(nil)

	ctrl := controllers.NewSerialsCtrl(mockRepo)
	err := ctrl.UpdateSerial(&models.Serial{S_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateSerial", &models.Serial{S_id: 1})
}

func TestDeleteSerial(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerials)
	mockRepo.On("DeleteSerial", 1).Return(nil)

	ctrl := controllers.NewSerialsCtrl(mockRepo)
	err := ctrl.DeleteSerial(1)

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteSerial", 1)
}

func TestGetSerialByTitle(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerials)
	mockRepo.On("GetSerialsByTitle", "title").Return([]*models.Serial{{S_id: 1}}, nil)

	ctrl := controllers.NewSerialsCtrl(mockRepo)
	serials, err := ctrl.GetSerialByTitle("title")

	require.NoError(t, err)
	assert.Equal(t, 1, serials[0].S_id)
	mockRepo.AssertCalled(t, "GetSerialsByTitle", "title")
}
