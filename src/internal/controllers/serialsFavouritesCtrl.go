package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type SerialsFavouritesCtrl struct {
	SerialsFavouritesService interfaces.IRepoSerialsFavourites
}

func NewSerialsFavouritesCtrl(service interfaces.IRepoSerialsFavourites) *SerialsFavouritesCtrl {
	return &SerialsFavouritesCtrl{SerialsFavouritesService: service}
}

func (ctrl *SerialsFavouritesCtrl) GetSerialsFavourites() ([]*models.SerialsFavourites, error) {
	return ctrl.SerialsFavouritesService.GetSerialsFavourites()
}

func (ctrl *SerialsFavouritesCtrl) GetSerialsByFavouriteId(id int) ([]*models.SerialsFavourites, error) {
	return ctrl.SerialsFavouritesService.GetSerialsByFavouriteId(id)
}

func (ctrl *SerialsFavouritesCtrl) GetFavouritesBySerialId(id int) ([]*models.SerialsFavourites, error) {
	return ctrl.SerialsFavouritesService.GetFavouritesBySerialId(id)
}

func (ctrl *SerialsFavouritesCtrl) GetSerialsFavouritesById(id int) (*models.SerialsFavourites, error) {
	return ctrl.SerialsFavouritesService.GetSerialsFavouritesById(id)
}

func (ctrl *SerialsFavouritesCtrl) CreateSerialsFavourites(serialFavourite *models.SerialsFavourites) error {
	return ctrl.SerialsFavouritesService.CreateSerialsFavourites(serialFavourite)
}

func (ctrl *SerialsFavouritesCtrl) UpdateSerialsFavourites(serialFavourite *models.SerialsFavourites) error {
	return ctrl.SerialsFavouritesService.UpdateSerialsFavourites(serialFavourite)
}

func (ctrl *SerialsFavouritesCtrl) CheckSerialInFavourite(serialFavourite *models.SerialsFavourites) bool {
	return ctrl.SerialsFavouritesService.CheckSerialInFavourite(serialFavourite)
}

func (ctrl *SerialsFavouritesCtrl) DeleteSerialById(idfav, idserial int) error {
	return ctrl.SerialsFavouritesService.DeleteSerialById(idfav, idserial)
}

func (ctrl *SerialsFavouritesCtrl) DeleteSerialsFavourites(id int) error {
	return ctrl.SerialsFavouritesService.DeleteSerialsFavourites(id)
}
