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

type FavouritesRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewFavouritesRepoMongo(client *mongo.Client, log *logrus.Logger) *FavouritesRepoMongo {
	db := client.Database("mydb")
	return &FavouritesRepoMongo{db: db, log: log}
}

func (repo *FavouritesRepoMongo) GetFavourites() ([]*models.Favourites, error) {
	repo.log.Info("Getting all favourites from the database")
	collection := repo.db.Collection("favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var favourites []*models.Favourites
	for cursor.Next(ctx) {
		var favourite models.Favourites
		if err := cursor.Decode(&favourite); err != nil {
			return nil, err
		}
		favourites = append(favourites, &favourite)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return favourites, nil
}

func (repo *FavouritesRepoMongo) GetFavouriteById(id int) (*models.Favourites, error) {
	repo.log.Info("Getting favourite by id from the database")
	collection := repo.db.Collection("favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var favourite models.Favourites
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&favourite)
	if err != nil {
		return nil, err
	}
	return &favourite, nil
}

func (repo *FavouritesRepoMongo) CreateFavourite(favourite *models.Favourites) (int, error) {
	if !favourite.Validate() {
		return 0, models.ErrInvalidModel
	}

	repo.log.Info("Creating favourite in the database")
	collection := repo.db.Collection("favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"f_cntSerials": favourite.GetCntSerials(),
	})
	if err != nil {
		return 0, err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	favourite.SetId(res)

	return favourite.GetId(), nil
}

func (repo *FavouritesRepoMongo) UpdateFavourite(favourite *models.Favourites) error {
	if !favourite.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating favourite in the database")
	collection := repo.db.Collection("favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(favourite.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"f_cntSerials": favourite.GetCntSerials(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FavouritesRepoMongo) DeleteFavourite(id int) error {
	repo.log.Info("Deleting favourite from the database")
	collection := repo.db.Collection("favourites")
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
