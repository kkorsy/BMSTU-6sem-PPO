package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type SerialsUsersCtrl struct {
	SerialsUsersService interfaces.IRepoSerialsUsers
}

func NewSerialsUsersCtrl(service interfaces.IRepoSerialsUsers) *SerialsUsersCtrl {
	return &SerialsUsersCtrl{SerialsUsersService: service}
}

func (ctrl *SerialsUsersCtrl) GetSerialsUsers() ([]*models.SerialsUsers, error) {
	return ctrl.SerialsUsersService.GetSerialsUsers()
}

func (ctrl *SerialsUsersCtrl) GetSerialsByUserId(id int) ([]*models.SerialsUsers, error) {
	return ctrl.SerialsUsersService.GetSerialsByUserId(id)
}

func (ctrl *SerialsUsersCtrl) GetUsersBySerialId(id int) ([]*models.SerialsUsers, error) {
	return ctrl.SerialsUsersService.GetUsersBySerialId(id)
}

func (ctrl *SerialsUsersCtrl) GetSerialsUsersById(id int) (*models.SerialsUsers, error) {
	return ctrl.SerialsUsersService.GetSerialsUsersById(id)
}

func (ctrl *SerialsUsersCtrl) GetSerialUserByIds(userId, serialId int) (*models.SerialsUsers, error) {
	return ctrl.SerialsUsersService.GetSerialUserByIds(userId, serialId)
}

func (ctrl *SerialsUsersCtrl) CreateSerialsUsers(serialUser *models.SerialsUsers) error {
	return ctrl.SerialsUsersService.CreateSerialsUsers(serialUser)
}

func (ctrl *SerialsUsersCtrl) UpdateSerialsUsers(serialUser *models.SerialsUsers) error {
	return ctrl.SerialsUsersService.UpdateSerialsUsers(serialUser)
}

func (ctrl *SerialsUsersCtrl) DeleteSerialsByUserId(id int) error {
	return ctrl.SerialsUsersService.DeleteSerialsByUserId(id)
}
