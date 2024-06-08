package postgres

import (
	"app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ActorsRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewActorsRepoPostgres(db *sqlx.DB, log *logrus.Logger) *ActorsRepoPostgres {
	return &ActorsRepoPostgres{db: db, log: log}
}

func (repo *ActorsRepoPostgres) GetActors() ([]*models.Actors, error) {
	repo.log.Info("Getting all actors from the database")
	actors := []*models.Actors{}
	err := repo.db.Select(&actors, "SELECT * FROM actors")
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (repo *ActorsRepoPostgres) GetActorById(id int) (*models.Actors, error) {
	repo.log.Info("Getting actor by id from the database")
	actor := &models.Actors{}
	err := repo.db.Get(actor, "SELECT * FROM actors WHERE a_id=$1", id)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (repo *ActorsRepoPostgres) CreateActor(actor *models.Actors) error {
	if !actor.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating actor in the database")
	err := repo.db.QueryRow("INSERT INTO actors (a_name, a_surname, a_gender, a_bdate) VALUES ($1, $2, $3, $4) RETURNING a_id",
		actor.GetName(), actor.GetSurname(), actor.GetGender(), actor.GetBdate()).Scan(&id)
	if err != nil {
		return err
	}
	actor.SetId(int(id))

	return nil
}

func (repo *ActorsRepoPostgres) UpdateActor(actor *models.Actors) error {
	if !actor.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating actor in the database")
	_, err := repo.db.Exec("UPDATE actors SET a_name=$1, a_surname=$2, a_gender=$3, a_bdate=$4 WHERE a_id=$5",
		actor.GetName(), actor.GetSurname(), actor.GetGender(), actor.GetBdate(), actor.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *ActorsRepoPostgres) DeleteActor(id int) error {
	repo.log.Info("Deleting actor from the database")
	_, err := repo.db.Exec("DELETE FROM actors WHERE a_id=$1", id)
	repo.log.Info(err)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorsRepoPostgres) CheckActor(actor *models.Actors) bool {
	repo.log.Info("Checking actor in the database")
	err := repo.db.Get(actor, "SELECT * FROM actors WHERE a_name=$1 AND a_surname=$2 AND a_gender=$3 AND a_bdate=$4", actor.GetName(), actor.GetSurname(), actor.GetGender(), actor.GetBdate())
	return err == nil
}
