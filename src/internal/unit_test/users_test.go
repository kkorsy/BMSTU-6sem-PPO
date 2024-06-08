package unit_test

import (
	"testing"

	"app/internal/controllers"
	"app/internal/mocks"
	"app/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestGetUsers(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	mockData := []*models.Users{
		{
			U_id:       1,
			U_login:    "test",
			U_password: "test",
		},
	}
	mockRepo.On("GetUsers").Return(mockData, nil)
	ctrl := controllers.NewUsersCtrl(mockRepo, nil)

	users, err := ctrl.GetUsers()
	require.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, mockData, users)
}

func TestGetUserById(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	mockData := &models.Users{
		U_id:       1,
		U_login:    "test",
		U_password: "test",
	}

	mockRepo.On("GetUserById", 1).Return(mockData, nil)

	user, err := mockRepo.GetUserById(1)
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, mockData, user)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	mockData := &models.Users{
		U_id:       1,
		U_login:    "test",
		U_password: "test",
	}

	mockRepo.On("CreateUser", mockData).Return(nil)

	err := mockRepo.CreateUser(mockData)
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	mockData := &models.Users{
		U_id:       1,
		U_login:    "test",
		U_password: "test",
	}

	mockRepo.On("UpdateUser", mockData).Return(nil)

	err := mockRepo.UpdateUser(mockData)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)

	mockRepo.On("DeleteUser", 1).Return(nil)

	err := mockRepo.DeleteUser(1)
	require.NoError(t, err)
}

func TestGetUserByLogin(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	mockData := &models.Users{
		U_id:       2,
		U_login:    "user",
		U_password: "password",
	}

	mockRepo.On("GetUserByLogin", "user").Return(mockData, nil)

	user, err := mockRepo.GetUserByLogin("user")
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, mockData.U_id, user.U_id)
	assert.Equal(t, mockData.U_login, user.U_login)
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(user.U_password), []byte(mockData.U_password)))
}

func TestAuthUser_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	ctrl := controllers.NewUsersCtrl(mockRepo, nil)

	mockRepo.On("AuthUser", "test", "password").Return(nil, controllers.ErrUserNotFound)

	usr, err := ctrl.AuthUser("test", "password")
	assert.EqualError(t, err, controllers.ErrUserNotFound.Error())
	assert.Nil(t, usr)
}

func TestAuthUser_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	ctrl := controllers.NewUsersCtrl(mockRepo, nil)

	mockRepo.On("AuthUser", "user", "wrongpassword").Return(nil, controllers.ErrInvalidPass)

	usr, err := ctrl.AuthUser("user", "wrongpassword")
	assert.EqualError(t, err, controllers.ErrInvalidPass.Error())
	assert.Nil(t, usr)
}

func TestAuthUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockRepoUsers)
	ctrl := controllers.NewUsersCtrl(mockRepo, nil)

	mockUser := &models.Users{
		U_id:       1,
		U_login:    "admin",
		U_password: "password",
	}

	mockRepo.On("AuthUser", "admin", "password").Return(mockUser, nil)

	usr, err := ctrl.AuthUser("admin", "password")
	require.NoError(t, err)
	assert.NotNil(t, usr)
}
