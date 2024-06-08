package interfaces

import (
	"app/internal/models"
)

type IRepoSerialsUsers interface {
	GetSerialsUsers() ([]*models.SerialsUsers, error)
	GetSerialsByUserId(id int) ([]*models.SerialsUsers, error)
	GetUsersBySerialId(id int) ([]*models.SerialsUsers, error)
	GetSerialsUsersById(id int) (*models.SerialsUsers, error)
	GetSerialUserByIds(userId, serialId int) (*models.SerialsUsers, error)
	CreateSerialsUsers(serialUser *models.SerialsUsers) error
	UpdateSerialsUsers(serialUser *models.SerialsUsers) error
	DeleteSerialsByUserId(id int) error
}
