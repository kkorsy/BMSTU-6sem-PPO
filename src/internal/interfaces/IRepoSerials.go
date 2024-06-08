package interfaces

import (
	"app/internal/models"
)

type IRepoSerials interface {
	GetSerials() ([]*models.Serial, error)
	GetSerialById(id int) (*models.Serial, error)
	GetSerialsByTitle(title string) ([]*models.Serial, error)
	CreateSerial(serial *models.Serial) error
	UpdateSerial(serial *models.Serial) error
	DeleteSerial(id int) error
	CalculateDuration(serial *models.Serial) error
}
