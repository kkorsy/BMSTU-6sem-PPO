package postgres

import (
	"app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProducersRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewProducersRepoPostgres(db *sqlx.DB, log *logrus.Logger) *ProducersRepoPostgres {
	return &ProducersRepoPostgres{db: db, log: log}
}

func (repo *ProducersRepoPostgres) GetProducers() ([]*models.Producers, error) {
	repo.log.Info("Getting all producers from the database")
	producers := []*models.Producers{}
	err := repo.db.Select(&producers, "SELECT * FROM producers")
	if err != nil {
		return nil, err
	}
	return producers, nil
}

func (repo *ProducersRepoPostgres) GetProducerById(id int) (*models.Producers, error) {
	repo.log.Info("Getting producer by id from the database")
	producer := &models.Producers{}
	err := repo.db.Get(producer, "SELECT * FROM producers WHERE p_id=$1", id)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (repo *ProducersRepoPostgres) CreateProducer(producer *models.Producers) error {
	if !producer.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating producer in the database")
	err := repo.db.QueryRow("INSERT INTO producers (p_name, p_surname) VALUES ($1, $2) RETURNING p_id",
		producer.GetName(), producer.GetSurname()).Scan(&id)
	if err != nil {
		return err
	}
	producer.SetId(int(id))

	return nil
}

func (repo *ProducersRepoPostgres) UpdateProducer(producer *models.Producers) error {
	if !producer.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating producer in the database")
	_, err := repo.db.Exec("UPDATE producers SET p_name=$1, p_surname=$2 WHERE p_id=$3",
		producer.GetName(), producer.GetSurname(), producer.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *ProducersRepoPostgres) DeleteProducer(id int) error {
	repo.log.Info("Deleting producer from the database")
	_, err := repo.db.Exec("DELETE FROM producers WHERE p_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
