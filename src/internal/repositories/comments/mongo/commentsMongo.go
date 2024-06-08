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

type CommentsRepoMongo struct {
	db  *mongo.Database
	log *logrus.Logger
}

func NewCommentsRepoMongo(client *mongo.Client, log *logrus.Logger) *CommentsRepoMongo {
	db := client.Database("mydb")
	return &CommentsRepoMongo{db: db, log: log}
}

func (repo *CommentsRepoMongo) GetComments() ([]*models.Comments, error) {
	repo.log.Info("Getting all comments from the database")
	comments := []*models.Comments{}
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comments
		if err = cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *CommentsRepoMongo) GetCommentById(id int) (*models.Comments, error) {
	repo.log.Info("Getting comment by id from the database")
	comment := &models.Comments{}
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (repo *CommentsRepoMongo) GetCommentsBySerialId(idSerial int) ([]*models.Comments, error) {
	repo.log.Info("Getting comments by serial from the database")
	comments := []*models.Comments{}
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"c_idSerial": idSerial})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comments
		if err = cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *CommentsRepoMongo) GetCommentsByUserId(idUser int) ([]*models.Comments, error) {
	repo.log.Info("Getting comments by user from the database")
	comments := []*models.Comments{}
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"c_idUser": idUser})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comments
		if err = cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *CommentsRepoMongo) GetCommentsBySerialIdUserId(idSerial, idUser int) (*models.Comments, error) {
	repo.log.Info("Getting comments by serial and user from the database")
	comment := &models.Comments{}
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"c_idSerial": idSerial, "c_idUser": idUser}).Decode(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (repo *CommentsRepoMongo) CreateComment(comment *models.Comments) error {
	if !comment.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Creating comment in the database")
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return err
	}
	res, _ := strconv.Atoi(result.InsertedID.(primitive.ObjectID).Hex())
	comment.SetId(res)

	return nil
}

func (repo *CommentsRepoMongo) UpdateComment(comment *models.Comments) error {
	if !comment.Validate() {
		return models.ErrInvalidModel
	}

	repo.log.Info("Updating comment in the database")
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(strconv.Itoa(comment.GetId()))
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"c_text":     comment.GetText(),
			"c_date":     comment.GetDate(),
			"c_idUser":   comment.GetIdUser(),
			"c_idSerial": comment.GetIdSerial(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CommentsRepoMongo) DeleteComment(id int) error {
	repo.log.Info("Deleting comment from the database")
	collection := repo.db.Collection("comments")
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

func (repo *CommentsRepoMongo) CheckComment(idUser, idSerial int) bool {
	repo.log.Info("Checking if comment exists in the database")
	collection := repo.db.Collection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"c_idUser":   idUser,
		"c_idSerial": idSerial,
	}

	err := collection.FindOne(ctx, filter).Err()
	return err == nil
}
