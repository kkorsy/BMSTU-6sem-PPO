package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/serialsFavourites/mongo"
	pg "app/internal/repositories/serialsFavourites/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSerialsFavouritesRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSerialsFavourites {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSerialsFavouritesRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSerialsFavouritesRepoMongo(db, log)
	default:
		return nil
	}
}
