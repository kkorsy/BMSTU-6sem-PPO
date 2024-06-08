package interfaces

import (
	"app/internal/models"
)

type IRepoEpisodes interface {
	GetEpisodes() ([]*models.Episodes, error)
	GetEpisodeById(id int) (*models.Episodes, error)
	GetEpisodesBySeasonId(id int) ([]*models.Episodes, error)
	CreateEpisode(episode *models.Episodes) error
	UpdateEpisode(episode *models.Episodes) error
	DeleteEpisode(id int) error
}
