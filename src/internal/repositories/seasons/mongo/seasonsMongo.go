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

type SeasonsRepo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewSeasonsRepoMongo(client *mongo.Client, log *logrus.Logger) *SeasonsRepo {
	db := client.Database("mydb")
	return &SeasonsRepo{db: db, log: log}
}

func (repo *SeasonsRepo) FormatDate(season *models.Seasons) {
	date := season.GetDate()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	season.SetDate(d2)
}

func (repo *SeasonsRepo) FormatDateList(seasons []*models.Seasons) {
	for _, season := range seasons {
		repo.FormatDate(season)
	}
}

func (repo *SeasonsRepo) GetSeasons() ([]*models.Seasons, error) {
	repo.log.Info("Getting all seasons from the database")
	collection := repo.db.Collection("seasons")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var seasons []*models.Seasons
	for cursor.Next(ctx) {
		var season models.Seasons
		if err := cursor.Decode(&season); err != nil {
			return nil, err
		}
		seasons = append(seasons, &season)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(seasons)
	return seasons, nil
}

func (repo *SeasonsRepo) GetSeasonById(id int) (*models.Seasons, error) {
	repo.log.Info("Getting season by id from the database")
	collection := repo.db.Collection("seasons")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var season models.Seasons
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&season)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(&season)
	return &season, nil
}

func (repo *SeasonsRepo) GetSeasonsBySerialId(id int) ([]*models.Seasons, error) {
	repo.log.Info("Getting seasons by serial id from the database")
	collection := repo.db.Collection("seasons")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(ctx, bson.M{"ss_idSerial": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var seasons []*models.Seasons
	for cursor.Next(ctx) {
		var season models.Seasons
		if err := cursor.Decode(&season); err != nil {
			return nil, err
		}
		seasons = append(seasons, &season)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(seasons)
	return seasons, nil
}

func (repo *SeasonsRepo) CreateSeason(season *models.Seasons) error {
	if !season.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating season in the database")
	collection := repo.db.Collection("seasons")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"ss_name":        season.GetName(),
		"ss_date":        season.GetDate(),
		"ss_idSerial":    season.GetIdSerial(),
		"ss_num":         season.GetNum(),
		"ss_cntEpisodes": season.GetCntEpisodes(),
	})
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	season.SetId(res)
	return nil
}

func (repo *SeasonsRepo) UpdateSeason(season *models.Seasons) error {
	if !season.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating season in the database")
	collection := repo.db.Collection("seasons")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(season.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"ss_name":        season.GetName(),
			"ss_date":        season.GetDate(),
			"ss_idSerial":    season.GetIdSerial(),
			"ss_num":         season.GetNum(),
			"ss_cntEpisodes": season.GetCntEpisodes(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SeasonsRepo) DeleteSeason(id int) error {
	repo.log.Info("Deleting season from the database")
	collection := repo.db.Collection("seasons")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
