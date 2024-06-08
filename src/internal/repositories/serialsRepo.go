package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/serials/mongo"
	pg "app/internal/repositories/serials/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerials {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsRepoMongo(db, log)
	default:
		return nil
	}
}
