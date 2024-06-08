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

type UsersRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewUsersRepoMongo(client *mongo.Client, log *logrus.Logger) *UsersRepoMongo {
	db := client.Database("mydb")
	return &UsersRepoMongo{db: db, log: log}
}

func (repo *UsersRepoMongo) FormatDate(user *models.Users) {
	date := user.GetBdate()
	d1, _ := time.Parse("2006-01-02T00:00:00Z", date)
	d2 := d1.Format("02.01.2006")
	user.SetBdate(d2)
}

func (repo *UsersRepoMongo) FormatDateList(users []*models.Users) {
	for _, user := range users {
		repo.FormatDate(user)
	}
}

func (repo *UsersRepoMongo) GetUsers() ([]*models.Users, error) {
	repo.log.Info("Getting all users from the database")
	collection := repo.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.Users
	for cursor.Next(ctx) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	repo.FormatDateList(users)
	return users, nil
}

func (repo *UsersRepoMongo) GetUserById(id int) (*models.Users, error) {
	repo.log.Info("Getting user by id from the database")
	collection := repo.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var user models.Users
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(&user)
	return &user, nil
}

func (repo *UsersRepoMongo) GetUserByLogin(login string) (*models.Users, error) {
	repo.log.Info("Getting user by login from the database")
	collection := repo.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user models.Users
	err := collection.FindOne(ctx, bson.M{"u_login": login}).Decode(&user)
	if err != nil {
		return nil, err
	}
	repo.FormatDate(&user)
	return &user, nil
}

func (repo *UsersRepoMongo) CreateUser(user *models.Users) error {
	if !user.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating user in the database")
	collection := repo.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"u_login":        user.GetLogin(),
		"u_password":     user.GetPassword(),
		"u_role":         user.GetRole(),
		"u_name":         user.GetName(),
		"u_surname":      user.GetSurname(),
		"u_gender":       user.GetGender(),
		"u_bdate":        user.GetBdate(),
		"u_idFavourites": user.GetIdFavourites(),
	})
	if err != nil {
		return err
	}

	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	user.SetId(res)
	return nil
}

func (repo *UsersRepoMongo) UpdateUser(user *models.Users) error {
	if !user.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating user in the database")
	collection := repo.db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(user.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"u_login":        user.GetLogin(),
			"u_password":     user.GetPassword(),
			"u_role":         user.GetRole(),
			"u_name":         user.GetName(),
			"u_surname":      user.GetSurname(),
			"u_gender":       user.GetGender(),
			"u_bdate":        user.GetBdate(),
			"u_idFavourites": user.GetIdFavourites(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		repo.log.Error(err)
		return err
	}
	return nil
}

func (repo *UsersRepoMongo) DeleteUser(id int) error {
	repo.log.Info("Deleting user from the database")
	collection := repo.db.Collection("users")
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

func (repo *UsersRepoMongo) CheckUser(login string) bool {
	repo.log.Info("Checking user by login from the database")
	_, err := repo.GetUserByLogin(login)
	return err == nil
}
