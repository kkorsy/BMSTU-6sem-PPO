package mocks

import (
	"app/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoComments struct {
	mock.Mock
}

func (m *MockRepoComments) GetComments() ([]*models.Comments, error) {
	args := m.Called()
	return args.Get(0).([]*models.Comments), args.Error(1)
}

func (m *MockRepoComments) GetCommentById(id int) (*models.Comments, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Comments), args.Error(1)
}

func (m *MockRepoComments) GetCommentsBySerialId(idSerial int) ([]*models.Comments, error) {
	args := m.Called(idSerial)
	return args.Get(0).([]*models.Comments), args.Error(1)
}

func (m *MockRepoComments) GetCommentsByUserId(idUser int) ([]*models.Comments, error) {
	args := m.Called(idUser)
	return args.Get(0).([]*models.Comments), args.Error(1)
}

func (m *MockRepoComments) GetCommentsBySerialIdUserId(idSerial, idUser int) (*models.Comments, error) {
	args := m.Called(idSerial, idUser)
	return args.Get(0).(*models.Comments), args.Error(1)
}

func (m *MockRepoComments) CreateComment(comment *models.Comments) error {
	args := m.Called(comment)
	return args.Error(0)
}

func (m *MockRepoComments) UpdateComment(comment *models.Comments) error {
	args := m.Called(comment)
	return args.Error(0)
}

func (m *MockRepoComments) DeleteComment(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepoComments) CheckComment(idUser, idSerial int) bool {
	args := m.Called(idUser, idSerial)
	return args.Bool(0)
}
