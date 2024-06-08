package interfaces

import "app/internal/models"

type IRepoStatistic interface {
	GetStatistic() (*models.Statistic, error)
	UpdateStatistic(stat *models.Statistic) error
}
