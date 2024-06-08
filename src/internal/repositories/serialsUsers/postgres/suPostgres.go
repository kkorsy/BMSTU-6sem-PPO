package postgres

import (
	"app/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SerialsUsersRepoPostgres struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewSerialsUsersRepoPostgres(db *sqlx.DB, log *logrus.Logger) *SerialsUsersRepoPostgres {
	return &SerialsUsersRepoPostgres{db: db, log: log}
}

func (repo *SerialsUsersRepoPostgres) FormatDate(su *models.SerialsUsers) {
	date := su.GetLastSeen()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	su.SetLastSeen(d2)
}

func (repo *SerialsUsersRepoPostgres) FormatDateList(suList []*models.SerialsUsers) {
	for _, su := range suList {
		repo.FormatDate(su)
	}
}

func (repo *SerialsUsersRepoPostgres) GetSerialsUsers() ([]*models.SerialsUsers, error) {
	repo.log.Info("Getting all serials_users from the database")
	serialsUsers := []*models.SerialsUsers{}
	err := repo.db.Select(&serialsUsers, "SELECT * FROM serials_users")
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(serialsUsers)
	return serialsUsers, nil
}

func (repo *SerialsUsersRepoPostgres) GetSerialsByUserId(id int) ([]*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by user id from the database")
	serialsUsers := []*models.SerialsUsers{}
	err := repo.db.Select(&serialsUsers, "SELECT * FROM serials_users WHERE su_idUser=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(serialsUsers)
	return serialsUsers, nil
}

func (repo *SerialsUsersRepoPostgres) GetUsersBySerialId(id int) ([]*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by serial id from the database")
	serialsUsers := []*models.SerialsUsers{}
	err := repo.db.Select(&serialsUsers, "SELECT * FROM serials_users WHERE su_idSerial=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDateList(serialsUsers)
	return serialsUsers, nil
}

func (repo *SerialsUsersRepoPostgres) GetSerialsUsersById(id int) (*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by id from the database")
	serialUser := &models.SerialsUsers{}
	err := repo.db.Get(serialUser, "SELECT * FROM serials_users WHERE su_id=$1", id)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(serialUser)
	return serialUser, nil
}

func (repo *SerialsUsersRepoPostgres) GetSerialUserByIds(serialId, userId int) (*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by user id and serial id from the database")
	serialUser := &models.SerialsUsers{}
	err := repo.db.Get(serialUser, "SELECT * FROM serials_users WHERE su_idUser=$1 AND su_idSerial=$2", userId, serialId)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(serialUser)
	return serialUser, nil
}

func (repo *SerialsUsersRepoPostgres) CreateSerialsUsers(serialUser *models.SerialsUsers) error {
	if !serialUser.Validate() {
		return models.ErrInvalidModel
	}
	var id int64

	repo.log.Info("Creating serials_users in the database")
	err := repo.db.QueryRow("INSERT INTO serials_users (su_idSerial, su_idUser, su_lastSeen) VALUES ($1, $2, $3) RETURNING su_id",
		serialUser.GetIdSerial(), serialUser.GetIdUser(), serialUser.GetLastSeen()).Scan(&id)
	if err != nil {
		return err
	}
	serialUser.SetId(int(id))

	return nil
}

func (repo *SerialsUsersRepoPostgres) UpdateSerialsUsers(serialUser *models.SerialsUsers) error {
	if !serialUser.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serials_users in the database")
	_, err := repo.db.Exec("UPDATE serials_users SET su_idSerial=$1, su_idUser=$2 WHERE su_id=$3",
		serialUser.GetIdSerial(), serialUser.GetIdUser(), serialUser.GetId())

	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsUsersRepoPostgres) DeleteSerialsByUserId(id int) error {
	repo.log.Info("Deleting serials_users by user id from the database")
	_, err := repo.db.Exec("DELETE FROM serials_users WHERE su_idUser=$1", id)
	if err != nil {
		return err
	}
	return nil
}
