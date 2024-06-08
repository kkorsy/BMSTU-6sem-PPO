package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type SeasonsCtrl struct {
	SeasonsService interfaces.IRepoSeasons
}

func NewSeasonsCtrl(service interfaces.IRepoSeasons) *SeasonsCtrl {
	return &SeasonsCtrl{SeasonsService: service}
}

func (ctrl *SeasonsCtrl) GetSeasons() ([]*models.Seasons, error) {
	return ctrl.SeasonsService.GetSeasons()
}

func (ctrl *SeasonsCtrl) GetSeasonById(id int) (*models.Seasons, error) {
	return ctrl.SeasonsService.GetSeasonById(id)
}

func (ctrl *SeasonsCtrl) GetSeasonsBySerialId(id int) ([]*models.Seasons, error) {
	return ctrl.SeasonsService.GetSeasonsBySerialId(id)
}

func (ctrl *SeasonsCtrl) CreateSeason(season *models.Seasons) error {
	return ctrl.SeasonsService.CreateSeason(season)
}

func (ctrl *SeasonsCtrl) UpdateSeason(season *models.Seasons) error {
	return ctrl.SeasonsService.UpdateSeason(season)
}

func (ctrl *SeasonsCtrl) DeleteSeason(id int) error {
	return ctrl.SeasonsService.DeleteSeason(id)
}
