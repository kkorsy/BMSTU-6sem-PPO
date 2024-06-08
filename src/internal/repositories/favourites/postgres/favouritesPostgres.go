package postgres

import (
	"app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type FavouritesRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewFavouritesRepoPostgres(db *sqlx.DB, log *logrus.Logger) *FavouritesRepoPostgres {
	return &FavouritesRepoPostgres{db: db, log: log}
}

func (repo *FavouritesRepoPostgres) GetFavourites() ([]*models.Favourites, error) {
	repo.log.Info("Getting all favourites from the database")
	favourites := []*models.Favourites{}
	err := repo.db.Select(&favourites, "SELECT * FROM favourites")
	if err != nil {
		return nil, err
	}
	return favourites, nil
}

func (repo *FavouritesRepoPostgres) GetFavouriteById(id int) (*models.Favourites, error) {
	repo.log.Info("Getting favourite by id from the database")
	favourite := &models.Favourites{}
	err := repo.db.Get(favourite, "SELECT * FROM favourites WHERE f_id=$1", id)
	if err != nil {
		return nil, err
	}
	return favourite, nil
}

func (repo *FavouritesRepoPostgres) CreateFavourite(favourite *models.Favourites) (int, error) {
	if !favourite.Validate() {
		return -1, models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating favourite in the database")
	err := repo.db.QueryRow("INSERT INTO favourites (f_cntSerials) VALUES ($1) RETURNING f_id",
		favourite.GetCntSerials()).Scan(&id)
	if err != nil {
		return -1, err
	}

	favourite.SetId(int(id))

	return favourite.GetId(), nil
}

func (repo *FavouritesRepoPostgres) UpdateFavourite(favourite *models.Favourites) error {
	if !favourite.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating favourite in the database")
	_, err := repo.db.Exec("UPDATE favourites SET f_cntSerials=$1 WHERE f_id=$2",
		favourite.GetCntSerials(), favourite.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *FavouritesRepoPostgres) DeleteFavourite(id int) error {
	repo.log.Info("Deleting favourite from the database")
	_, err := repo.db.Exec("DELETE FROM favourites WHERE f_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
