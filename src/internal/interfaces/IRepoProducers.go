package interfaces

import (
	"app/internal/models"
)

type IRepoProducers interface {
	GetProducers() ([]*models.Producers, error)
	GetProducerById(id int) (*models.Producers, error)
	CreateProducer(producer *models.Producers) error
	UpdateProducer(producer *models.Producers) error
	DeleteProducer(id int) error
}
