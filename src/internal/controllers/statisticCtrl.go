package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type StatisticCtrl struct {
	StatisticService interfaces.IRepoStatistic
}

func NewStatisticCtrl(Stservice interfaces.IRepoStatistic) *StatisticCtrl {
	return &StatisticCtrl{StatisticService: Stservice}
}

func (ctrl *StatisticCtrl) GetStatistic() (*models.Statistic, error) {
	return ctrl.StatisticService.GetStatistic()
}

func (ctrl *StatisticCtrl) UpdateStatistic(stat *models.Statistic) error {
	return ctrl.StatisticService.UpdateStatistic(stat)
}
