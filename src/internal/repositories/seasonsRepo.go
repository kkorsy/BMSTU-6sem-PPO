package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/seasons/mongo"
	pg "app/internal/repositories/seasons/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSeasonsRepo(db interface{}, log *logrus.Logger) interfaces.IRepoSeasons {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewSeasonsRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewSeasonsRepoMongo(db, log)
	default:
		return nil
	}
}
