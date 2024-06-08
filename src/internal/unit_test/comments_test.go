package unit_test

import (
	"testing"

	"app/internal/controllers"
	"app/internal/mocks"
	"app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetComments(t *testing.T) {
	mockRepo := new(mocks.MockRepoComments)
	mockRepo.On("GetComments").Return([]*models.Comments{{C_id: 1}}, nil)

	ctrl := controllers.NewCommentsCtrl(mockRepo)
	comments, err := ctrl.GetComments()

	require.NoError(t, err)
	assert.Equal(t, 1, comments[0].C_id)
	mockRepo.AssertCalled(t, "GetComments")
}

func TestGetCommentById(t *testing.T) {
	mockRepo := new(mocks.MockRepoComments)
	mockRepo.On("GetCommentById", 1).Return(&models.Comments{C_id: 1}, nil)

	ctrl := controllers.NewCommentsCtrl(mockRepo)
	comment, err := ctrl.GetCommentById(1)

	require.NoError(t, err)
	assert.Equal(t, 1, comment.C_id)
	mockRepo.AssertCalled(t, "GetCommentById", 1)
}

func TestCreateComment(t *testing.T) {
	mockRepo := new(mocks.MockRepoComments)
	mockRepo.On("CreateComment", &models.Comments{C_id: 1}).Return(nil)

	ctrl := controllers.NewCommentsCtrl(mockRepo)
	err := ctrl.CreateComment(&models.Comments{C_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateComment", &models.Comments{C_id: 1})
}

func TestUpdateComment(t *testing.T) {
	mockRepo := new(mocks.MockRepoComments)
	mockRepo.On("UpdateComment", &models.Comments{C_id: 1}).Return(nil)

	ctrl := controllers.NewCommentsCtrl(mockRepo)
	err := ctrl.UpdateComment(&models.Comments{C_id: 1})

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateComment", &models.Comments{C_id: 1})
}

func TestDeleteComment(t *testing.T) {
	mockRepo := new(mocks.MockRepoComments)
	mockRepo.On("DeleteComment", 1).Return(nil)

	ctrl := controllers.NewCommentsCtrl(mockRepo)
	err := ctrl.DeleteComment(1)

	require.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteComment", 1)
}
