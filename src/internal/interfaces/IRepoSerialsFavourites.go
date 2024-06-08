package interfaces

import (
	"app/internal/models"
)

type IRepoSerialsFavourites interface {
	GetSerialsFavourites() ([]*models.SerialsFavourites, error)
	GetSerialsByFavouriteId(id int) ([]*models.SerialsFavourites, error)
	GetFavouritesBySerialId(id int) ([]*models.SerialsFavourites, error)
	GetSerialsFavouritesById(id int) (*models.SerialsFavourites, error)
	CreateSerialsFavourites(serialFavourite *models.SerialsFavourites) error
	UpdateSerialsFavourites(serialFavourite *models.SerialsFavourites) error
	CheckSerialInFavourite(serialFavourite *models.SerialsFavourites) bool
	DeleteSerialById(idfav, idserial int) error
	DeleteSerialsFavourites(id int) error
}
