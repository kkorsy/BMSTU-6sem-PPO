package mocks

import (
	"app/internal/controllers"
	"app/internal/models"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockRepoUsers struct {
	mock.Mock
}

func (m *MockRepoUsers) GetUsers() ([]*models.Users, error) {
	args := m.Called()
	return args.Get(0).([]*models.Users), args.Error(1)
}

func (m *MockRepoUsers) GetUserById(id int) (*models.Users, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *MockRepoUsers) CheckUser(user string) bool {
	args := m.Called(user)
	return args.Bool(0)
}

func (m *MockRepoUsers) CreateUser(user *models.Users) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepoUsers) UpdateUser(user *models.Users) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepoUsers) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepoUsers) GetUserByLogin(login string) (*models.Users, error) {
	pass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if login == "admin" {
		return &models.Users{
			U_id:       1,
			U_login:    "admin",
			U_password: string(pass),
		}, nil
	} else if login == "user" {
		return &models.Users{
			U_id:       2,
			U_login:    "user",
			U_password: string(pass),
		}, nil
	}

	return nil, controllers.ErrUserNotFound
}
