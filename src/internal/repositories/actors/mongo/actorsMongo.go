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

type ActorsRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewActorsRepoMongo(client *mongo.Client, log *logrus.Logger) *ActorsRepoMongo {
	db := client.Database("mydb")
	return &ActorsRepoMongo{db: db, log: log}
}

func (repo *ActorsRepoMongo) GetActors() ([]*models.Actors, error) {
	repo.log.Info("Getting all actors from the database")
	actors := []*models.Actors{}
	collection := repo.db.Collection("actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var actor models.Actors
		if err = cursor.Decode(&actor); err != nil {
			return nil, err
		}
		actors = append(actors, &actor)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return actors, nil
}

func (repo *ActorsRepoMongo) GetActorById(id int) (*models.Actors, error) {
	repo.log.Info("Getting actor by id from the database")
	actor := &models.Actors{}
	collection := repo.db.Collection("actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(actor)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (repo *ActorsRepoMongo) CreateActor(actor *models.Actors) error {
	if !actor.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating actor in the database")
	collection := repo.db.Collection("actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, actor)
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	actor.SetId(res)

	return nil
}

func (repo *ActorsRepoMongo) UpdateActor(actor *models.Actors) error {
	if !actor.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating actor in the database")
	collection := repo.db.Collection("actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(actor.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"a_name":    actor.GetName(),
			"a_surname": actor.GetSurname(),
			"a_gender":  actor.GetGender(),
			"a_bdate":   actor.GetBdate(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ActorsRepoMongo) DeleteActor(id int) error {
	repo.log.Info("Deleting actor from the database")
	collection := repo.db.Collection("actors")
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

func (repo *ActorsRepoMongo) CheckActor(actor *models.Actors) bool {
	repo.log.Info("Checking actor in the database")
	collection := repo.db.Collection("actors")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"a_name":    actor.GetName(),
		"a_surname": actor.GetSurname(),
		"a_gender":  actor.GetGender(),
		"a_bdate":   actor.GetBdate(),
	}

	err := collection.FindOne(ctx, filter).Decode(actor)
	return err == nil
}
