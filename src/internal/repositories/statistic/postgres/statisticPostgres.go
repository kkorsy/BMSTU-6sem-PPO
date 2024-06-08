package postgres

import (
	"app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type StatisticRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewStatisticRepoPostgres(db *sqlx.DB, log *logrus.Logger) *StatisticRepoPostgres {
	return &StatisticRepoPostgres{db: db, log: log}
}

func (repo *StatisticRepoPostgres) GetStatistic() (*models.Statistic, error) {
	repo.log.Info("Getting statistic from the database")
	stat := &models.Statistic{}
	err := repo.db.Get(stat, "SELECT * FROM statistic")
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func (repo *StatisticRepoPostgres) UpdateStatistic(stat *models.Statistic) error {
	if !stat.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating statistic in the database")
	_, err := repo.db.Exec("UPDATE statistic SET st_gender_male=$1, st_gender_female=$2, st_role_user=$3, st_role_admin=$4, st_age_0_18=$5, st_age_19_30=$6, st_age_31_50=$7, st_age_51_100=$8 WHERE st_id=$9",
		stat.GetGenderMale(), stat.GetGenderFemale(), stat.GetRoleUser(), stat.GetRoleAdmin(), stat.GetAge0_18(), stat.GetAge19_30(), stat.GetAge31_50(), stat.GetAge51_100(), stat.GetId())
	if err != nil {
		return err
	}
	return nil
}
