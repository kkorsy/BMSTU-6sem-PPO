package postgres

import (
	"app/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type EpisodesRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewEpisodesRepoPostgres(db *sqlx.DB, log *logrus.Logger) *EpisodesRepoPostgres {
	return &EpisodesRepoPostgres{db: db, log: log}
}

func (repo *EpisodesRepoPostgres) FormatDate(episode *models.Episodes) {
	date := episode.GetDate()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	episode.SetDate(d2)
}

func (repo *EpisodesRepoPostgres) FormatDateList(episodes []*models.Episodes) {
	for _, episode := range episodes {
		repo.FormatDate(episode)
	}
}

func (repo *EpisodesRepoPostgres) GetEpisodes() ([]*models.Episodes, error) {
	repo.log.Info("Getting all episodes from the database")
	episodes := []*models.Episodes{}
	err := repo.db.Select(&episodes, "SELECT * FROM episodes")
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(episodes)
	return episodes, nil
}

func (repo *EpisodesRepoPostgres) GetEpisodeById(id int) (*models.Episodes, error) {
	repo.log.Info("Getting episode by id from the database")
	episode := &models.Episodes{}
	err := repo.db.Get(episode, "SELECT * FROM episodes WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(episode)
	return episode, nil
}

func (repo *EpisodesRepoPostgres) GetEpisodesBySeasonId(id int) ([]*models.Episodes, error) {
	repo.log.Info("Getting episodes by season id from the database")
	episodes := []*models.Episodes{}
	err := repo.db.Select(&episodes, "SELECT * FROM episodes WHERE e_idSeason=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(episodes)
	return episodes, nil
}

func (repo *EpisodesRepoPostgres) CreateEpisode(episode *models.Episodes) error {
	if !episode.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating episode in the database")
	err := repo.db.QueryRow("INSERT INTO episodes (e_name, e_date, e_idSeason, e_num, e_duration) VALUES ($1, $2, $3, $4, $5) RETURNING e_id",
		episode.GetName(), episode.GetDate(), episode.GetIdSeason(), episode.GetNum(), episode.GetDuration()).Scan(&id)
	if err != nil {
		return err
	}
	episode.SetId(int(id))

	return nil
}

func (repo *EpisodesRepoPostgres) UpdateEpisode(episode *models.Episodes) error {
	if !episode.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating episode in the database")
	_, err := repo.db.Exec("UPDATE episodes SET e_name=$1, e_date=$2, e_idSeason=$3, e_num=$4, e_duration=$5 WHERE e_id=$6",
		episode.GetName(), episode.GetDate(), episode.GetIdSeason(), episode.GetNum(), episode.GetDuration(), episode.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *EpisodesRepoPostgres) DeleteEpisode(id int) error {
	repo.log.Info("Deleting episode from the database")
	_, err := repo.db.Exec("DELETE FROM episodes WHERE e_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
