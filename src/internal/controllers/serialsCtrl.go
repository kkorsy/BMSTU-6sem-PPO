package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type SerialsCtrl struct {
	SerialsService interfaces.IRepoSerials
	SeasonsService interfaces.IRepoSeasons
}

func NewSerialsCtrl(service interfaces.IRepoSerials) *SerialsCtrl {
	return &SerialsCtrl{SerialsService: service}
}

func (ctrl *SerialsCtrl) GetSerials() ([]*models.Serial, error) {
	return ctrl.SerialsService.GetSerials()
}

func (ctrl *SerialsCtrl) GetSerialById(id int) (*models.Serial, error) {
	return ctrl.SerialsService.GetSerialById(id)
}

func (ctrl *SerialsCtrl) CreateSerial(serial *models.Serial) error {
	err := ctrl.SerialsService.CalculateDuration(serial)
	if err != nil {
		return err
	}
	return ctrl.SerialsService.CreateSerial(serial)
}

func (ctrl *SerialsCtrl) UpdateSerial(serial *models.Serial) error {
	return ctrl.SerialsService.UpdateSerial(serial)
}

func (ctrl *SerialsCtrl) DeleteSerial(id int) error {
	return ctrl.SerialsService.DeleteSerial(id)
}

func (ctrl *SerialsCtrl) GetSerialByTitle(title string) ([]*models.Serial, error) {
	return ctrl.SerialsService.GetSerialsByTitle(title)
}

func (ctrl *SerialsCtrl) CountSeasons(id int) (int, error) {
	seasons, err := ctrl.SeasonsService.GetSeasonsBySerialId(id)
	if err != nil {
		return 0, err
	}
	return len(seasons), nil
}
