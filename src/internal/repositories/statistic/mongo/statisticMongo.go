package mongo

import (
	"app/internal/models"
	"context"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatisticRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewStatisticRepoMongo(client *mongo.Client, log *logrus.Logger) *StatisticRepoMongo {
	db := client.Database("mydb")
	return &StatisticRepoMongo{db: db, log: log}
}

func (repo *StatisticRepoMongo) GetStatistic() (*models.Statistic, error) {
	repo.log.Info("Getting statistic from the database")
	collection := repo.db.Collection("statistic")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stat := &models.Statistic{}
	err := collection.FindOne(ctx, bson.M{}).Decode(stat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return stat, nil
}

func (repo *StatisticRepoMongo) UpdateStatistic(stat *models.Statistic) error {
	if !stat.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating statistic in the database")
	collection := repo.db.Collection("statistic")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(stat.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"st_gender_male":   stat.GetGenderMale(),
			"st_gender_female": stat.GetGenderFemale(),
			"st_role_user":     stat.GetRoleUser(),
			"st_role_admin":    stat.GetRoleAdmin(),
			"st_age_0_18":      stat.GetAge0_18(),
			"st_age_19_30":     stat.GetAge19_30(),
			"st_age_31_50":     stat.GetAge31_50(),
			"st_age_51_100":    stat.GetAge51_100(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}
