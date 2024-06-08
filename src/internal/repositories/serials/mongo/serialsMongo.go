package mongo

import (
	"app/internal/models"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SerialsRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewSerialsRepoMongo(client *mongo.Client, log *logrus.Logger) *SerialsRepoMongo {
	db := client.Database("mydb")
	return &SerialsRepoMongo{db: db, log: log}
}

func (repo *SerialsRepoMongo) GetSerials() ([]*models.Serial, error) {
	repo.log.Info("Getting all serials from the database")
	serials := []*models.Serial{}
	collection := repo.db.Collection("serials")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var serial models.Serial
		if err = cursor.Decode(&serial); err != nil {
			return nil, err
		}
		serials = append(serials, &serial)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return serials, nil
}

func (repo *SerialsRepoMongo) GetSerialById(id int) (*models.Serial, error) {
	repo.log.Info("Getting serial by id from the database")
	serial := &models.Serial{}
	collection := repo.db.Collection("serials")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(serial)
	if err != nil {
		return nil, err
	}
	return serial, nil
}

func (repo *SerialsRepoMongo) GetSerialsByTitle(title string) ([]*models.Serial, error) {
	repo.log.Info("Getting serial by title from the database")
	serials := []*models.Serial{}
	collection := repo.db.Collection("serials")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"s_name": bson.M{"$regex": title, "$options": "i"}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var serial models.Serial
		if err = cursor.Decode(&serial); err != nil {
			return nil, err
		}
		serials = append(serials, &serial)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return serials, nil
}

func (repo *SerialsRepoMongo) CreateSerial(serial *models.Serial) error {
	if !serial.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating serial in the database")
	collection := repo.db.Collection("serials")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, serial)
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	serial.SetId(res)

	return nil
}

func (repo *SerialsRepoMongo) UpdateSerial(serial *models.Serial) error {
	if !serial.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serial in the database")
	collection := repo.db.Collection("serials")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(serial.GetId()))
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{
			"s_idProducer":  serial.GetIdProducer(),
			"s_name":        serial.GetName(),
			"s_description": serial.GetDescription(),
			"s_year":        serial.GetYear(),
			"s_genre":       serial.GetGenre(),
			"s_rating":      serial.GetRating(),
			"s_seasons":     serial.GetSeasons(),
			"s_state":       serial.GetState(),
			"s_duration":    serial.S_duration,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsRepoMongo) DeleteSerial(id int) error {
	repo.log.Info("Deleting serial from the database")
	collection := repo.db.Collection("serials")
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

func (repo *SerialsRepoMongo) CalculateDuration(serial *models.Serial) error {
	repo.log.Info("Calculating duration of the serial")
	collection := repo.db.Collection("serials")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", serial.GetId()}}}},
		{{"$group", bson.D{
			{"_id", "$_id"},
			{"totalDuration", bson.D{{"$sum", "$s_duration"}}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result struct {
			TotalDuration string `bson:"totalDuration"`
		}
		if err := cursor.Decode(&result); err != nil {
			return err
		}
		serial.S_duration = result.TotalDuration
	} else {
		return fmt.Errorf("no result found for serial ID %d", serial.GetId())
	}

	return nil
}
