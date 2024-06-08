package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type ProducersCtrl struct {
	ProducersService interfaces.IRepoProducers
}

func NewProducersCtrl(service interfaces.IRepoProducers) *ProducersCtrl {
	return &ProducersCtrl{ProducersService: service}
}

func (ctrl *ProducersCtrl) GetProducers() ([]*models.Producers, error) {
	return ctrl.ProducersService.GetProducers()
}

func (ctrl *ProducersCtrl) GetProducerById(id int) (*models.Producers, error) {
	return ctrl.ProducersService.GetProducerById(id)
}

func (ctrl *ProducersCtrl) CreateProducer(producer *models.Producers) error {
	return ctrl.ProducersService.CreateProducer(producer)
}

func (ctrl *ProducersCtrl) UpdateProducer(producer *models.Producers) error {
	return ctrl.ProducersService.UpdateProducer(producer)
}

func (ctrl *ProducersCtrl) DeleteProducer(id int) error {
	return ctrl.ProducersService.DeleteProducer(id)
}
