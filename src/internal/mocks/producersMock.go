package mocks

import (
	"app/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockRepoProducers struct {
	mock.Mock
}

func (m *MockRepoProducers) GetProducers() ([]*models.Producers, error) {
	args := m.Called()
	return args.Get(0).([]*models.Producers), args.Error(1)
}

func (m *MockRepoProducers) GetProducerById(id int) (*models.Producers, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Producers), args.Error(1)
}

func (m *MockRepoProducers) CreateProducer(producer *models.Producers) error {
	args := m.Called(producer)
	return args.Error(0)
}

func (m *MockRepoProducers) UpdateProducer(producer *models.Producers) error {
	args := m.Called(producer)
	return args.Error(0)
}

func (m *MockRepoProducers) DeleteProducer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
