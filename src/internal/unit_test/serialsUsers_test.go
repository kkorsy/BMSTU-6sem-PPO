package unit_test

import (
	"testing"

	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSerialsUsers(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsUsers)
	mockData := []*models.SerialsUsers{
		{
			Su_id:       1,
			Su_idSerial: 1,
			Su_idUser:   1,
		},
	}

	mockRepo.On("GetSerialsUsers").Return(mockData, nil)

	serialsUsers, err := mockRepo.GetSerialsUsers()
	require.NoError(t, err)
	assert.NotNil(t, serialsUsers)
	assert.Equal(t, mockData, serialsUsers)
}

func TestGetSerialsUsersById(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsUsers)
	mockData := &models.SerialsUsers{
		Su_id:       1,
		Su_idSerial: 1,
		Su_idUser:   1,
	}

	mockRepo.On("GetSerialsUsersById", 1).Return(mockData, nil)

	serialsUsers, err := mockRepo.GetSerialsUsersById(1)
	require.NoError(t, err)
	assert.NotNil(t, serialsUsers)
	assert.Equal(t, mockData, serialsUsers)
}

func TestGetSerialsByUserId(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsUsers)
	mockData := []*models.SerialsUsers{
		{
			Su_id:       1,
			Su_idSerial: 1,
			Su_idUser:   1,
		},
	}

	mockRepo.On("GetSerialsByUserId", 1).Return(mockData, nil)

	serialsUsers, err := mockRepo.GetSerialsByUserId(1)
	require.NoError(t, err)
	assert.NotNil(t, serialsUsers)
	assert.Equal(t, mockData, serialsUsers)
}

func TestGetUsersBySerialId(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsUsers)
	mockData := []*models.SerialsUsers{
		{
			Su_id:       1,
			Su_idSerial: 1,
			Su_idUser:   1,
		},
	}

	mockRepo.On("GetUsersBySerialId", 1).Return(mockData, nil)

	serialsUsers, err := mockRepo.GetUsersBySerialId(1)
	require.NoError(t, err)
	assert.NotNil(t, serialsUsers)
	assert.Equal(t, mockData, serialsUsers)
}

func TestCreateSerialsUsers(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsUsers)
	mockData := &models.SerialsUsers{
		Su_id:       1,
		Su_idSerial: 1,
		Su_idUser:   1,
	}

	mockRepo.On("CreateSerialsUsers", mockData).Return(nil)

	err := mockRepo.CreateSerialsUsers(mockData)
	require.NoError(t, err)
}

func TestUpdateSerialsUsers(t *testing.T) {
	mockRepo := new(mocks.MockRepoSerialsUsers)
	mockData := &models.SerialsUsers{
		Su_id:       1,
		Su_idSerial: 1,
		Su_idUser:   1,
	}

	mockRepo.On("UpdateSerialsUsers", mockData).Return(nil)

	err := mockRepo.UpdateSerialsUsers(mockData)
	require.NoError(t, err)
}
