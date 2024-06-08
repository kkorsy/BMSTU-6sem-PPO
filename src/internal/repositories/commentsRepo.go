package repositories

import (
	"app/internal/interfaces"
	mg "app/internal/repositories/comments/mongo"
	pg "app/internal/repositories/comments/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCommentsRepo(db interface{}, log *logrus.Logger) interfaces.IRepoComments {
	switch db := db.(type) {
	case *sqlx.DB:
		return pg.NewCommentsRepoPostgres(db, log)
	case *mongo.Client:
		return mg.NewCommentsRepoMongo(db, log)
	default:
		return nil
	}
}
