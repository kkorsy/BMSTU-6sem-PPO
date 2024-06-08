package interfaces

import (
	"app/internal/models"
)

type IRepoSerialsActors interface {
	GetSerialsActors() ([]*models.SerialsActors, error)
	GetSerialsByActorId(id int) ([]*models.SerialsActors, error)
	GetActorsBySerialId(id int) ([]*models.SerialsActors, error)
	GetSerialsActorsById(id int) (*models.SerialsActors, error)
	CreateSerialsActors(serialActor *models.SerialsActors) error
	UpdateSerialsActors(serialActor *models.SerialsActors) error
	DeleteSerialsActors(id int) error
}
