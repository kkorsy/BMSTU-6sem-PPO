package interfaces

import (
	"app/internal/models"
)

type IRepoComments interface {
	GetComments() ([]*models.Comments, error)
	GetCommentById(id int) (*models.Comments, error)
	GetCommentsBySerialId(idSerial int) ([]*models.Comments, error)
	GetCommentsByUserId(idUser int) ([]*models.Comments, error)
	GetCommentsBySerialIdUserId(idSerial, idUser int) (*models.Comments, error)
	CreateComment(comment *models.Comments) error
	UpdateComment(comment *models.Comments) error
	DeleteComment(id int) error
	CheckComment(idUser, idSerial int) bool
}
