package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/serialsUsers/mongo"
	pg "app/internal/repositories/serialsUsers/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsUsersRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerialsUsers {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsUsersRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsUsersRepoMongo(db, log)
	default:
		return nil
	}
}
