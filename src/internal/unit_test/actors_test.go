package unit_test

import (
	"testing"

	"app/internal/controllers"
	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetActors(t *testing.T) {
	mockRepo := new(mocks.MockRepoActors)
	mockRepo.On("GetActors").Return([]*models.Actors{{A_id: 1}}, nil)

	ctrl := controllers.NewActorsCtrl(mockRepo)
	actors, err := ctrl.GetActors()

	require.NoError(t, err)
	assert.Equal(t, 1, actors[0].A_id)
	mockRepo.AssertCalled(t, "GetActors")
}

func TestGetActorById(t *testing.T) {
	mockRepo := new(mocks.MockRepoActors)
	mockRepo.On("GetActorById", 1).Return(&models.Actors{A_id: 1}, nil)

	ctrl := controllers.NewActorsCtrl(mockRepo)
	actor, err := ctrl.GetActorById(1)

	require.NoError(t, err)
	assert.Equal(t, 1, actor.A_id)
	mockRepo.AssertCalled(t, "GetActorById", 1)
}

func TestCreateActor(t *testing.T) {
	mockRepo := new(mocks.MockRepoActors)
	mockRepo.On("CreateActor", &models.Actors{A_id: 1}).Return(nil)

	ctrl := controllers.NewActorsCtrl(mockRepo)
	err := ctrl.CreateActor(&models.Actors{A_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateActor", &models.Actors{A_id: 1})
}

func TestUpdateActor(t *testing.T) {
	mockRepo := new(mocks.MockRepoActors)
	mockRepo.On("UpdateActor", &models.Actors{A_id: 1}).Return(nil)

	ctrl := controllers.NewActorsCtrl(mockRepo)
	err := ctrl.UpdateActor(&models.Actors{A_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateActor", &models.Actors{A_id: 1})
}

func TestDeleteActor(t *testing.T) {
	mockRepo := new(mocks.MockRepoActors)
	mockRepo.On("DeleteActor", 1).Return(nil)

	ctrl := controllers.NewActorsCtrl(mockRepo)
	err := ctrl.DeleteActor(1)

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteActor", 1)
}
