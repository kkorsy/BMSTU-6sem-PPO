package postgres

import (
	"app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SerialsRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewSerialsRepoPostgres(db *sqlx.DB, log *logrus.Logger) *SerialsRepoPostgres {
	return &SerialsRepoPostgres{db: db, log: log}
}

func (repo *SerialsRepoPostgres) GetSerials() ([]*models.Serial, error) {
	repo.log.Info("Getting all serials from the database")
	serials := []*models.Serial{}
	err := repo.db.Select(&serials, "SELECT * FROM serials")
	if err != nil {
		return nil, err
	}
	return serials, nil
}

func (repo *SerialsRepoPostgres) GetSerialById(id int) (*models.Serial, error) {
	repo.log.Info("Getting serial by id from the database")
	serial := &models.Serial{}
	err := repo.db.Get(serial, "SELECT * FROM serials WHERE s_id=$1", id)
	if err != nil {
		return nil, err
	}
	return serial, nil
}

func (repo *SerialsRepoPostgres) GetSerialsByTitle(title string) ([]*models.Serial, error) {
	repo.log.Info("Getting serial by title from the database")
	serials := []*models.Serial{}
	err := repo.db.Select(&serials, "SELECT * FROM serials WHERE s_name LIKE $1", string("%"+title+"%"))
	if err != nil {
		return nil, err
	}
	return serials, nil
}

func (repo *SerialsRepoPostgres) CreateSerial(serial *models.Serial) error {
	if !serial.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating serial in the database")
	err := repo.db.QueryRow("INSERT INTO serials (s_idProducer, s_name, s_description, s_year, s_genre, s_rating, s_seasons, s_state, s_img, s_duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING s_id",
		serial.GetIdProducer(), serial.GetName(), serial.GetDescription(), serial.GetYear(), serial.GetGenre(), serial.GetRating(), serial.GetSeasons(), serial.GetState(), serial.S_img, serial.S_duration).Scan(&id)
	if err != nil {
		return err
	}
	serial.SetId(int(id))

	return nil
}

func (repo *SerialsRepoPostgres) UpdateSerial(serial *models.Serial) error {
	if !serial.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serial in the database")
	_, err := repo.db.Exec("UPDATE serials SET s_idProducer=$2, s_name=$3, s_description=$4, s_year=$5, s_genre=$6, s_rating=$7, s_seasons=$8, s_state=$9, s_duration=$10 WHERE s_id=$1",
		serial.GetId(), serial.GetIdProducer(), serial.GetName(), serial.GetDescription(), serial.GetYear(), serial.GetGenre(), serial.GetRating(), serial.GetSeasons(), serial.GetState(), serial.S_duration)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsRepoPostgres) DeleteSerial(id int) error {
	repo.log.Info("Deleting serial from the database")
	_, err := repo.db.Exec("DELETE FROM serials WHERE s_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SerialsRepoPostgres) CalculateDuration(serial *models.Serial) error {
	repo.log.Info("Calculating duration of the serial")
	err := repo.db.QueryRow("SELECT calculate_total_duration($1)", serial.GetId()).Scan(&serial.S_duration)
	if err != nil {
		return err
	}
	return nil
}
