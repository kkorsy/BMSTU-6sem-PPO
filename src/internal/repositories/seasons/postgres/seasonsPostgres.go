package postgres

import (
	"app/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SeasonsRepo struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewSeasonsRepoPostgres(db *sqlx.DB, log *logrus.Logger) *SeasonsRepo {
	return &SeasonsRepo{db: db, log: log}
}

func (repo *SeasonsRepo) FormatDate(season *models.Seasons) {
	date := season.GetDate()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	season.SetDate(d2)
}

func (repo *SeasonsRepo) FormatDateList(seasons []*models.Seasons) {
	for _, season := range seasons {
		repo.FormatDate(season)
	}
}

func (repo *SeasonsRepo) GetSeasons() ([]*models.Seasons, error) {
	repo.log.Info("Getting all seasons from the database")
	seasons := []*models.Seasons{}
	err := repo.db.Select(&seasons, "SELECT * FROM seasons")
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(seasons)
	return seasons, nil
}

func (repo *SeasonsRepo) GetSeasonById(id int) (*models.Seasons, error) {
	repo.log.Info("Getting season by id from the database")
	season := &models.Seasons{}
	err := repo.db.Get(season, "SELECT * FROM seasons WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(season)
	return season, nil
}

func (repo *SeasonsRepo) GetSeasonsBySerialId(id int) ([]*models.Seasons, error) {
	repo.log.Info("Getting seasons by serial id from the database")
	seasons := []*models.Seasons{}
	err := repo.db.Select(&seasons, "SELECT * FROM seasons WHERE ss_idSerial=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(seasons)
	return seasons, nil
}

func (repo *SeasonsRepo) CreateSeason(season *models.Seasons) error {
	if !season.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating season in the database")
	err := repo.db.QueryRow("INSERT INTO seasons (ss_id, ss_name, ss_date, ss_idSerial, ss_num, ss_cntEpisodes) VALUES ($1, $2, $3, $4, $5) RETURNING ss_id",
		season.GetName(), season.GetDate(), season.GetIdSerial(), season.GetNum(), season.GetCntEpisodes()).Scan(&id)
	if err != nil {
		return err
	}
	season.SetId(int(id))

	return nil
}

func (repo *SeasonsRepo) UpdateSeason(season *models.Seasons) error {
	if !season.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating season in the database")
	_, err := repo.db.Exec("UPDATE seasons SET ss_name=$1, ss_date=$2, ss_idSerial=$3, ss_num=$4, ss_cntEpisodes=$5 WHERE ss_id=$6",
		season.GetName(), season.GetDate(), season.GetIdSerial(), season.GetNum(), season.GetCntEpisodes(), season.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *SeasonsRepo) DeleteSeason(id int) error {
	repo.log.Info("Deleting season from the database")
	_, err := repo.db.Exec("DELETE FROM seasons WHERE s_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
