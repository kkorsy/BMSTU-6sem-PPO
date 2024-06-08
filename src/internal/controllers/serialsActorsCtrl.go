package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type SerialsActorsCtrl struct {
	SerialsActorsService interfaces.IRepoSerialsActors
}

func NewSerialsActorsCtrl(service interfaces.IRepoSerialsActors) *SerialsActorsCtrl {
	return &SerialsActorsCtrl{SerialsActorsService: service}
}

func (ctrl *SerialsActorsCtrl) GetSerialsActors() ([]*models.SerialsActors, error) {
	return ctrl.SerialsActorsService.GetSerialsActors()
}

func (ctrl *SerialsActorsCtrl) GetSerialsByActorId(id int) ([]*models.SerialsActors, error) {
	return ctrl.SerialsActorsService.GetSerialsByActorId(id)
}

func (ctrl *SerialsActorsCtrl) GetActorsBySerialId(id int) ([]*models.SerialsActors, error) {
	return ctrl.SerialsActorsService.GetActorsBySerialId(id)
}

func (ctrl *SerialsActorsCtrl) GetSerialsActorsById(id int) (*models.SerialsActors, error) {
	return ctrl.SerialsActorsService.GetSerialsActorsById(id)
}

func (ctrl *SerialsActorsCtrl) CreateSerialsActors(serialActor *models.SerialsActors) error {
	return ctrl.SerialsActorsService.CreateSerialsActors(serialActor)
}

func (ctrl *SerialsActorsCtrl) UpdateSerialsActors(serialActor *models.SerialsActors) error {
	return ctrl.SerialsActorsService.UpdateSerialsActors(serialActor)
}

func (ctrl *SerialsActorsCtrl) DeleteSerialsActors(id int) error {
	return ctrl.SerialsActorsService.DeleteSerialsActors(id)
}
