package postgres

import (
	"app/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UsersRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewUsersRepoPostgres(db *sqlx.DB, log *logrus.Logger) *UsersRepoPostgres {
	return &UsersRepoPostgres{db: db, log: log}
}

func (repo *UsersRepoPostgres) FormatDate(user *models.Users) {
	date := user.GetBdate()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	user.SetBdate(d2)
}

func (repo *UsersRepoPostgres) FormatDateList(users []*models.Users) {
	for _, user := range users {
		repo.FormatDate(user)
	}
}

func (repo *UsersRepoPostgres) GetUsers() ([]*models.Users, error) {
	repo.log.Info("Getting all users from the database")
	users := []*models.Users{}
	err := repo.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(users)
	return users, nil
}

func (repo *UsersRepoPostgres) GetUserById(id int) (*models.Users, error) {
	repo.log.Info("Getting user by id from the database")
	user := &models.Users{}
	err := repo.db.Get(user, "SELECT * FROM users WHERE u_id=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(user)
	return user, nil
}

func (repo *UsersRepoPostgres) GetUserByLogin(login string) (*models.Users, error) {
	repo.log.Info("Getting user by login from the database")
	user := &models.Users{}
	err := repo.db.Get(user, "SELECT * FROM users WHERE u_login=$1", login)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(user)
	return user, nil
}

func (repo *UsersRepoPostgres) CreateUser(user *models.Users) error {
	if !user.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating user in the database")
	err := repo.db.QueryRow("INSERT INTO users (u_login, u_password, u_role, u_name, u_surname, u_gender, u_bdate, u_idFavourites) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING u_id",
		user.GetLogin(), user.GetPassword(), user.GetRole(), user.GetName(), user.GetSurname(), user.GetGender(), user.GetBdate(), user.GetIdFavourites()).Scan(&id)
	if err != nil {
		return err
	}
	user.SetId(int(id))

	return nil
}

func (repo *UsersRepoPostgres) UpdateUser(user *models.Users) error {
	if !user.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating user in the database")
	_, err := repo.db.Exec("UPDATE users SET u_login=$1, u_password=$2, u_role=$3, u_name=$4, u_surname=$5, u_gender=$6, u_bdate=$7, u_idFavourites=$8 WHERE u_id=$9",
		user.GetLogin(), user.GetPassword(), user.GetRole(), user.GetName(), user.GetSurname(), user.GetGender(), user.GetBdate(), user.GetIdFavourites(), user.GetId())

	if err != nil {
		repo.log.Error(err)
		return err
	}
	return nil
}

func (repo *UsersRepoPostgres) DeleteUser(id int) error {
	repo.log.Info("Deleting user from the database")
	_, err := repo.db.Exec("DELETE FROM users WHERE u_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UsersRepoPostgres) CheckUser(login string) bool {
	repo.log.Info("Checking user by login from the database")
	_, err := repo.GetUserByLogin(login)
	return err == nil
}
