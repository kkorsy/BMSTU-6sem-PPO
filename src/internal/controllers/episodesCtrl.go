package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type EpisodesCtrl struct {
	EpisodesService interfaces.IRepoEpisodes
}

func NewEpisodesCtrl(service interfaces.IRepoEpisodes) *EpisodesCtrl {
	return &EpisodesCtrl{EpisodesService: service}
}

func (ctrl *EpisodesCtrl) GetEpisodes() ([]*models.Episodes, error) {
	return ctrl.EpisodesService.GetEpisodes()
}

func (ctrl *EpisodesCtrl) GetEpisodeById(id int) (*models.Episodes, error) {
	return ctrl.EpisodesService.GetEpisodeById(id)
}

func (ctrl *EpisodesCtrl) GetEpisodesBySeasonId(id int) ([]*models.Episodes, error) {
	return ctrl.EpisodesService.GetEpisodesBySeasonId(id)
}

func (ctrl *EpisodesCtrl) CreateEpisode(episode *models.Episodes) error {
	return ctrl.EpisodesService.CreateEpisode(episode)
}

func (ctrl *EpisodesCtrl) UpdateEpisode(episode *models.Episodes) error {
	return ctrl.EpisodesService.UpdateEpisode(episode)
}

func (ctrl *EpisodesCtrl) DeleteEpisode(id int) error {
	return ctrl.EpisodesService.DeleteEpisode(id)
}
