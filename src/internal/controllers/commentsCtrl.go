package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"
)

type CommentsCtrl struct {
	CommentsService interfaces.IRepoComments
}

func NewCommentsCtrl(service interfaces.IRepoComments) *CommentsCtrl {
	return &CommentsCtrl{CommentsService: service}
}

func (ctrl *CommentsCtrl) GetComments() ([]*models.Comments, error) {
	return ctrl.CommentsService.GetComments()
}

func (ctrl *CommentsCtrl) GetCommentById(id int) (*models.Comments, error) {
	return ctrl.CommentsService.GetCommentById(id)
}

func (ctrl *CommentsCtrl) GetCommentsBySerialId(idSerial int) ([]*models.Comments, error) {
	return ctrl.CommentsService.GetCommentsBySerialId(idSerial)
}

func (ctrl *CommentsCtrl) GetCommentsByUserId(idUser int) ([]*models.Comments, error) {
	return ctrl.CommentsService.GetCommentsByUserId(idUser)
}

func (ctrl *CommentsCtrl) GetCommentsBySerialIdUserId(idSerial, idUser int) (*models.Comments, error) {
	return ctrl.CommentsService.GetCommentsBySerialIdUserId(idSerial, idUser)
}

func (ctrl *CommentsCtrl) CreateComment(comment *models.Comments) error {
	return ctrl.CommentsService.CreateComment(comment)
}

func (ctrl *CommentsCtrl) UpdateComment(comment *models.Comments) error {
	return ctrl.CommentsService.UpdateComment(comment)
}

func (ctrl *CommentsCtrl) DeleteComment(id int) error {
	return ctrl.CommentsService.DeleteComment(id)
}

func (ctrl *CommentsCtrl) CheckComment(idUser, idSerial int) bool {
	return ctrl.CommentsService.CheckComment(idUser, idSerial)
}
