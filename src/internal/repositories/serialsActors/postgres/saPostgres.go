package postgres

import (
	"app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SerialsActorsRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewSerialsActorsRepoPostgres(db *sqlx.DB, log *logrus.Logger) *SerialsActorsRepoPostgres {
	return &SerialsActorsRepoPostgres{db: db, log: log}
}

func (repo *SerialsActorsRepoPostgres) GetSerialsActors() ([]*models.SerialsActors, error) {
	repo.log.Info("Getting all serials_actors from the database")
	serialsActors := []*models.SerialsActors{}
	err := repo.db.Select(&serialsActors, "SELECT * FROM serials_actors")
	if err != nil {
		return nil, err
	}
	return serialsActors, nil
}

func (repo *SerialsActorsRepoPostgres) GetSerialsActorsById(id int) (*models.SerialsActors, error) {
	repo.log.Info("Getting serials_actors by id from the database")
	serialActor := &models.SerialsActors{}
	err := repo.db.Get(serialActor, "SELECT * FROM serials_actors WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return serialActor, nil
}

func (repo *SerialsActorsRepoPostgres) CreateSerialsActors(serialActor *models.SerialsActors) error {
	if !serialActor.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating serials_actors in the database")
	err := repo.db.QueryRow("INSERT INTO serials_actors (sa_idSerial, sa_idActor) VALUES ($1, $2) RETURNING sa_id",
		serialActor.GetIdSerial(), serialActor.GetIdActor()).Scan(&id)
	if err != nil {
		return err
	}
	serialActor.SetId(int(id))

	return nil
}

func (repo *SerialsActorsRepoPostgres) UpdateSerialsActors(serialActor *models.SerialsActors) error {
	if !serialActor.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serials_actors in the database")
	_, err := repo.db.Exec("UPDATE serials_actors SET sa_idSerial=$1, sa_idActor=$2 WHERE sa_id=$3",
		serialActor.GetIdSerial(), serialActor.GetIdActor(), serialActor.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsActorsRepoPostgres) GetSerialsByActorId(id int) ([]*models.SerialsActors, error) {
	repo.log.Info("Getting serials by actor id from the database")
	serialsActors := []*models.SerialsActors{}
	err := repo.db.Select(&serialsActors, "SELECT * FROM serials_actors WHERE sa_idActor=$1", id)
	if err != nil {
		return nil, err
	}
	return serialsActors, nil
}

func (repo *SerialsActorsRepoPostgres) GetActorsBySerialId(id int) ([]*models.SerialsActors, error) {
	repo.log.Info("Getting actors by serial id from the database")
	serialsActors := []*models.SerialsActors{}
	err := repo.db.Select(&serialsActors, "SELECT * FROM serials_actors WHERE sa_idSerial=$1", id)
	if err != nil {
		return nil, err
	}
	return serialsActors, nil
}

func (repo *SerialsActorsRepoPostgres) DeleteSerialsActors(id int) error {
	repo.log.Info("Deleting serials_actors from the database")
	_, err := repo.db.Exec("DELETE FROM serials_actors WHERE sa_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
