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

type ProducersRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewProducersRepoMongo(client *mongo.Client, log *logrus.Logger) *ProducersRepoMongo {
	db := client.Database("mydb")
	return &ProducersRepoMongo{db: db, log: log}
}

func (repo *ProducersRepoMongo) GetProducers() ([]*models.Producers, error) {
	repo.log.Info("Getting all producers from the database")
	collection := repo.db.Collection("producers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var producers []*models.Producers
	for cursor.Next(ctx) {
		var producer models.Producers
		if err := cursor.Decode(&producer); err != nil {
			return nil, err
		}
		producers = append(producers, &producer)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return producers, nil
}

func (repo *ProducersRepoMongo) GetProducerById(id int) (*models.Producers, error) {
	repo.log.Info("Getting producer by id from the database")
	collection := repo.db.Collection("producers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var producer models.Producers
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&producer)
	if err != nil {
		return nil, err
	}
	return &producer, nil
}

func (repo *ProducersRepoMongo) CreateProducer(producer *models.Producers) error {
	if !producer.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating producer in the database")
	collection := repo.db.Collection("producers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"p_name":    producer.GetName(),
		"p_surname": producer.GetSurname(),
	})
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	producer.SetId(res)

	return nil
}

func (repo *ProducersRepoMongo) UpdateProducer(producer *models.Producers) error {
	if !producer.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating producer in the database")
	collection := repo.db.Collection("producers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(producer.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"p_name":    producer.GetName(),
			"p_surname": producer.GetSurname(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProducersRepoMongo) DeleteProducer(id int) error {
	repo.log.Info("Deleting producer from the database")
	collection := repo.db.Collection("producers")
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
