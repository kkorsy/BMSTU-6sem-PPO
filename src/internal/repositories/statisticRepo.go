package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/statistic/mongo"
	pg "app/internal/repositories/statistic/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewStatisticRepo(db interface{}, log *logrus.Logger) interfaces.IRepoStatistic {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewStatisticRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewStatisticRepoMongo(db, log)
	default:
		return nil
	}
}
