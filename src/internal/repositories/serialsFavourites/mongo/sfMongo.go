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

type SerialsFavouritesRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewSerialsFavouritesRepoMongo(client *mongo.Client, log *logrus.Logger) *SerialsFavouritesRepoMongo {
	db := client.Database("mydb")
	return &SerialsFavouritesRepoMongo{db: db, log: log}
}

func (repo *SerialsFavouritesRepoMongo) GetSerialsFavourites() ([]*models.SerialsFavourites, error) {
	repo.log.Info("Getting all serials_favourites from the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsFavourites []*models.SerialsFavourites
	for cursor.Next(ctx) {
		var serialFavourite models.SerialsFavourites
		if err := cursor.Decode(&serialFavourite); err != nil {
			return nil, err
		}
		serialsFavourites = append(serialsFavourites, &serialFavourite)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return serialsFavourites, nil
}

func (repo *SerialsFavouritesRepoMongo) GetSerialsFavouritesById(id int) (*models.SerialsFavourites, error) {
	repo.log.Info("Getting serials_favourites by id from the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var serialFavourite models.SerialsFavourites
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&serialFavourite)
	if err != nil {
		return nil, err
	}
	return &serialFavourite, nil
}

func (repo *SerialsFavouritesRepoMongo) GetSerialsByFavouriteId(id int) ([]*models.SerialsFavourites, error) {
	repo.log.Info("Getting serials_favourites by favourite id from the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"sf_idFavourite": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsFavourites []*models.SerialsFavourites
	for cursor.Next(ctx) {
		var serialFavourite models.SerialsFavourites
		if err := cursor.Decode(&serialFavourite); err != nil {
			return nil, err
		}
		serialsFavourites = append(serialsFavourites, &serialFavourite)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return serialsFavourites, nil
}

func (repo *SerialsFavouritesRepoMongo) GetFavouritesBySerialId(id int) ([]*models.SerialsFavourites, error) {
	repo.log.Info("Getting serials_favourites by serial id from the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"sf_idSerial": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsFavourites []*models.SerialsFavourites
	for cursor.Next(ctx) {
		var serialFavourite models.SerialsFavourites
		if err := cursor.Decode(&serialFavourite); err != nil {
			return nil, err
		}
		serialsFavourites = append(serialsFavourites, &serialFavourite)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return serialsFavourites, nil
}

func (repo *SerialsFavouritesRepoMongo) CreateSerialsFavourites(serialFavourite *models.SerialsFavourites) error {
	if !serialFavourite.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating serials_favourites in the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"sf_idSerial":    serialFavourite.GetIdSerial(),
		"sf_idFavourite": serialFavourite.GetIdFavourite(),
	})
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	serialFavourite.SetId(res)
	return nil
}

func (repo *SerialsFavouritesRepoMongo) UpdateSerialsFavourites(serialFavourite *models.SerialsFavourites) error {
	if !serialFavourite.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serials_favourites in the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(serialFavourite.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"sf_idSerial":    serialFavourite.GetIdSerial(),
			"sf_idFavourite": serialFavourite.GetIdFavourite(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsFavouritesRepoMongo) CheckSerialInFavourite(serialFavourite *models.SerialsFavourites) bool {
	repo.log.Info("Checking serial in favourite")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"sf_idSerial":    serialFavourite.GetIdSerial(),
		"sf_idFavourite": serialFavourite.GetIdFavourite(),
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false
	}

	return count > 0
}

func (repo *SerialsFavouritesRepoMongo) DeleteSerialById(idfav, idserial int) error {
	repo.log.Info("Deleting serials_favourites from the database")
	collection := repo.db.Collection("serials_favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"sf_idFavourite": idfav, "sf_idSerial": idserial})
	if err != nil {
		return err
	}
	return nil
}

func (repo *SerialsFavouritesRepoMongo) DeleteSerialsFavourites(id int) error {
	repo.log.Info("Deleting serials_favourites from the database")
	collection := repo.db.Collection("serials_favourites")
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
