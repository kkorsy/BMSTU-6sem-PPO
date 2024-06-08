package interfaces

import (
	"app/internal/models"
)

type IRepoFavourites interface {
	GetFavourites() ([]*models.Favourites, error)
	GetFavouriteById(id int) (*models.Favourites, error)
	CreateFavourite(favourite *models.Favourites) (int, error)
	UpdateFavourite(favourite *models.Favourites) error
	DeleteFavourite(id int) error
}
