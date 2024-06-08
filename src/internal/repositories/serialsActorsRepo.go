package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/serialsActors/mongo"
	pg "app/internal/repositories/serialsActors/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsActorsRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerialsActors {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsActorsRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsActorsRepoMongo(db, log)
	default:
		return nil
	}
}
