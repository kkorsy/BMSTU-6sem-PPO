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

type SerialsActorsRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewSerialsActorsRepoMongo(client *mongo.Client, log *logrus.Logger) *SerialsActorsRepoMongo {
	db := client.Database("mydb")
	return &SerialsActorsRepoMongo{db: db, log: log}
}

func (repo *SerialsActorsRepoMongo) GetSerialsActors() ([]*models.SerialsActors, error) {
	repo.log.Info("Getting all serials_actors from the database")
	collection := repo.db.Collection("serials_actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsActors []*models.SerialsActors
	for cursor.Next(ctx) {
		var serialActor models.SerialsActors
		if err := cursor.Decode(&serialActor); err != nil {
			return nil, err
		}
		serialsActors = append(serialsActors, &serialActor)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return serialsActors, nil
}

func (repo *SerialsActorsRepoMongo) GetSerialsActorsById(id int) (*models.SerialsActors, error) {
	repo.log.Info("Getting serials_actors by id from the database")
	collection := repo.db.Collection("serials_actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var serialActor models.SerialsActors
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&serialActor)
	if err != nil {
		return nil, err
	}
	return &serialActor, nil
}

func (repo *SerialsActorsRepoMongo) CreateSerialsActors(serialActor *models.SerialsActors) error {
	if !serialActor.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating serials_actors in the database")
	collection := repo.db.Collection("serials_actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"sa_idSerial": serialActor.GetIdSerial(),
		"sa_idActor":  serialActor.GetIdActor(),
	})
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	serialActor.SetId(res)
	return nil
}

func (repo *SerialsActorsRepoMongo) UpdateSerialsActors(serialActor *models.SerialsActors) error {
	if !serialActor.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serials_actors in the database")
	collection := repo.db.Collection("serials_actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(serialActor.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"sa_idSerial": serialActor.GetIdSerial(),
			"sa_idActor":  serialActor.GetIdActor(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsActorsRepoMongo) GetSerialsByActorId(id int) ([]*models.SerialsActors, error) {
	repo.log.Info("Getting serials by actor id from the database")
	collection := repo.db.Collection("serials_actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"sa_idActor": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsActors []*models.SerialsActors
	for cursor.Next(ctx) {
		var serialActor models.SerialsActors
		if err := cursor.Decode(&serialActor); err != nil {
			return nil, err
		}
		serialsActors = append(serialsActors, &serialActor)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return serialsActors, nil
}

func (repo *SerialsActorsRepoMongo) GetActorsBySerialId(id int) ([]*models.SerialsActors, error) {
	repo.log.Info("Getting actors by serial id from the database")
	collection := repo.db.Collection("serials_actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"sa_idSerial": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsActors []*models.SerialsActors
	for cursor.Next(ctx) {
		var serialActor models.SerialsActors
		if err := cursor.Decode(&serialActor); err != nil {
			return nil, err
		}
		serialsActors = append(serialsActors, &serialActor)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return serialsActors, nil
}

func (repo *SerialsActorsRepoMongo) DeleteSerialsActors(id int) error {
	repo.log.Info("Deleting serials_actors from the database")
	collection := repo.db.Collection("serials_actors")
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
