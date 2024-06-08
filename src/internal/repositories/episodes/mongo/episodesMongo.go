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

type EpisodesRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewEpisodesRepoMongo(client *mongo.Client, log *logrus.Logger) *EpisodesRepoMongo {
	db := client.Database("mydb")
	return &EpisodesRepoMongo{db: db, log: log}
}

func (repo *EpisodesRepoMongo) FormatDate(episode *models.Episodes) {
	date := episode.GetDate()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	episode.SetDate(d2)
}

func (repo *EpisodesRepoMongo) FormatDateList(episodes []*models.Episodes) {
	for _, episode := range episodes {
		repo.FormatDate(episode)
	}
}

func (repo *EpisodesRepoMongo) GetEpisodes() ([]*models.Episodes, error) {
	repo.log.Info("Getting all episodes from the database")
	episodes := []*models.Episodes{}
	collection := repo.db.Collection("episodes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var episode models.Episodes
		if err = cursor.Decode(&episode); err != nil {
			return nil, err
		}
		episodes = append(episodes, &episode)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(episodes)
	return episodes, nil
}

func (repo *EpisodesRepoMongo) GetEpisodeById(id int) (*models.Episodes, error) {
	repo.log.Info("Getting episode by id from the database")
	episode := &models.Episodes{}
	collection := repo.db.Collection("episodes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(episode)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(episode)
	return episode, nil
}

func (repo *EpisodesRepoMongo) GetEpisodesBySeasonId(idSeason int) ([]*models.Episodes, error) {
	repo.log.Info("Getting episodes by season id from the database")
	episodes := []*models.Episodes{}
	collection := repo.db.Collection("episodes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"e_idSeason": idSeason})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var episode models.Episodes
		if err = cursor.Decode(&episode); err != nil {
			return nil, err
		}
		episodes = append(episodes, &episode)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(episodes)
	return episodes, nil
}

func (repo *EpisodesRepoMongo) CreateEpisode(episode *models.Episodes) error {
	if !episode.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating episode in the database")
	collection := repo.db.Collection("episodes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"e_name":     episode.GetName(),
		"e_date":     episode.GetDate(),
		"e_idSeason": episode.GetIdSeason(),
		"e_num":      episode.GetNum(),
		"e_duration": episode.GetDuration(),
	})
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	episode.SetId(res)

	return nil
}

func (repo *EpisodesRepoMongo) UpdateEpisode(episode *models.Episodes) error {
	if !episode.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating episode in the database")
	collection := repo.db.Collection("episodes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(episode.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"e_name":     episode.GetName(),
			"e_date":     episode.GetDate(),
			"e_idSeason": episode.GetIdSeason(),
			"e_num":      episode.GetNum(),
			"e_duration": episode.GetDuration(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *EpisodesRepoMongo) DeleteEpisode(id int) error {
	repo.log.Info("Deleting episode from the database")
	collection := repo.db.Collection("episodes")
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
