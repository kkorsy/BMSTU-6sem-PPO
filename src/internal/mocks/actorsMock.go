package mocks

import (
	"app/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepoActors struct {
	mock.Mock
}

func (m *MockRepoActors) GetActors() ([]*models.Actors, error) {
	args := m.Called()
	return args.Get(0).([]*models.Actors), args.Error(1)
}

func (m *MockRepoActors) GetActorById(id int) (*models.Actors, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Actors), args.Error(1)
}

func (m *MockRepoActors) CreateActor(actor *models.Actors) error {
	args := m.Called(actor)
	return args.Error(0)
}

func (m *MockRepoActors) UpdateActor(actor *models.Actors) error {
	args := m.Called(actor)
	return args.Error(0)
}

func (m *MockRepoActors) DeleteActor(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepoActors) CheckActor(actor *models.Actors) bool {
	args := m.Called(actor)
	return args.Bool(0)
}
