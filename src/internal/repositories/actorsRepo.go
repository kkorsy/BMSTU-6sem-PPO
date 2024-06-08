package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/actors/mongo"
	pg "app/internal/repositories/actors/postgres"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewActorsRepo(db interface{}, log *logrus.Logger) interfaces.IRepoActors {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewActorsRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewActorsRepoMongo(db, log)
	default:
		return nil
	}
}
