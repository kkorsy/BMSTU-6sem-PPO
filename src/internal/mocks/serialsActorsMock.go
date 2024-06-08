package mocks

import (
	"app/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockRepoSerialsActors struct {
	mock.Mock
}

func (m *MockRepoSerialsActors) GetSerialsActors() ([]*models.SerialsActors, error) {
	args := m.Called()
	return args.Get(0).([]*models.SerialsActors), args.Error(1)
}

func (m *MockRepoSerialsActors) GetSerialsByActorId(id int) ([]*models.SerialsActors, error) {
	args := m.Called(id)
	return args.Get(0).([]*models.SerialsActors), args.Error(1)
}

func (m *MockRepoSerialsActors) GetActorsBySerialId(id int) ([]*models.SerialsActors, error) {
	args := m.Called(id)
	return args.Get(0).([]*models.SerialsActors), args.Error(1)
}

func (m *MockRepoSerialsActors) GetSerialsActorsById(id int) (*models.SerialsActors, error) {
	args := m.Called(id)
	return args.Get(0).(*models.SerialsActors), args.Error(1)
}

func (m *MockRepoSerialsActors) CreateSerialsActors(serialActor *models.SerialsActors) error {
	args := m.Called(serialActor)
	return args.Error(0)
}

func (m *MockRepoSerialsActors) UpdateSerialsActors(serialActor *models.SerialsActors) error {
	args := m.Called(serialActor)
	return args.Error(0)
}
