package interfaces

import (
	"app/internal/models"
)

type IRepoSeasons interface {
	GetSeasons() ([]*models.Seasons, error)
	GetSeasonById(id int) (*models.Seasons, error)
	GetSeasonsBySerialId(id int) ([]*models.Seasons, error)
	CreateSeason(season *models.Seasons) error
	UpdateSeason(season *models.Seasons) error
	DeleteSeason(id int) error
}
