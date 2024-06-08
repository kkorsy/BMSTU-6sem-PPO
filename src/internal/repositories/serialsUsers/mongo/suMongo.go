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

type SerialsUsersRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewSerialsUsersRepoMongo(client *mongo.Client, log *logrus.Logger) *SerialsUsersRepoMongo {
	db := client.Database("mydb")
	return &SerialsUsersRepoMongo{db: db, log: log}
}

func (repo *SerialsUsersRepoMongo) FormatDate(su *models.SerialsUsers) {
	date := su.GetLastSeen()
	d1, _ := time.Parse(time.RFC3339, date)
	d2 := d1.Format("02.01.2006")
	su.SetLastSeen(d2)
}

func (repo *SerialsUsersRepoMongo) FormatDateList(suList []*models.SerialsUsers) {
	for _, su := range suList {
		repo.FormatDate(su)
	}
}

func (repo *SerialsUsersRepoMongo) GetSerialsUsers() ([]*models.SerialsUsers, error) {
	repo.log.Info("Getting all serials_users from the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsUsers []*models.SerialsUsers
	for cursor.Next(ctx) {
		var serialUser models.SerialsUsers
		if err := cursor.Decode(&serialUser); err != nil {
			return nil, err
		}
		serialsUsers = append(serialsUsers, &serialUser)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(serialsUsers)
	return serialsUsers, nil
}

func (repo *SerialsUsersRepoMongo) GetSerialsByUserId(id int) ([]*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by user id from the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"su_idUser": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsUsers []*models.SerialsUsers
	for cursor.Next(ctx) {
		var serialUser models.SerialsUsers
		if err := cursor.Decode(&serialUser); err != nil {
			return nil, err
		}
		serialsUsers = append(serialsUsers, &serialUser)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(serialsUsers)
	return serialsUsers, nil
}

func (repo *SerialsUsersRepoMongo) GetUsersBySerialId(id int) ([]*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by serial id from the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"su_idSerial": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serialsUsers []*models.SerialsUsers
	for cursor.Next(ctx) {
		var serialUser models.SerialsUsers
		if err := cursor.Decode(&serialUser); err != nil {
			return nil, err
		}
		serialsUsers = append(serialsUsers, &serialUser)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(serialsUsers)
	return serialsUsers, nil
}

func (repo *SerialsUsersRepoMongo) GetSerialsUsersById(id int) (*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by id from the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var serialUser models.SerialsUsers
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&serialUser)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(&serialUser)
	return &serialUser, nil
}

func (repo *SerialsUsersRepoMongo) GetSerialUserByIds(serialId, userId int) (*models.SerialsUsers, error) {
	repo.log.Info("Getting serials_users by user id and serial id from the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var serialUser models.SerialsUsers
	err := collection.FindOne(ctx, bson.M{"su_idUser": userId, "su_idSerial": serialId}).Decode(&serialUser)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(&serialUser)
	return &serialUser, nil
}

func (repo *SerialsUsersRepoMongo) CreateSerialsUsers(serialUser *models.SerialsUsers) error {
	if !serialUser.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating serials_users in the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"su_idSerial": serialUser.GetIdSerial(),
		"su_idUser":   serialUser.GetIdUser(),
		"su_lastSeen": serialUser.GetLastSeen(),
	})
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	serialUser.SetId(res)
	return nil
}

func (repo *SerialsUsersRepoMongo) UpdateSerialsUsers(serialUser *models.SerialsUsers) error {
	if !serialUser.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating serials_users in the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(serialUser.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"su_idSerial": serialUser.GetIdSerial(),
			"su_idUser":   serialUser.GetIdUser(),
			"su_lastSeen": serialUser.GetLastSeen(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SerialsUsersRepoMongo) DeleteSerialsByUserId(id int) error {
	repo.log.Info("Deleting serials_users by user id from the database")
	collection := repo.db.Collection("serials_users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.DeleteMany(ctx, bson.M{"su_idUser": id})
	if err != nil {
		return err
	}
	return nil
}
