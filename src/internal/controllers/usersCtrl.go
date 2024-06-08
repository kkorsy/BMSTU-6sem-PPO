package controllers

import (
	"app/internal/interfaces"
	"app/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UsersCtrl struct {
	UsersService interfaces.IRepoUsers
	FavService   interfaces.IRepoFavourites
}

func NewUsersCtrl(Uservice interfaces.IRepoUsers, Fservice interfaces.IRepoFavourites) *UsersCtrl {
	return &UsersCtrl{UsersService: Uservice, FavService: Fservice}
}

func (ctrl *UsersCtrl) GetUsers() ([]*models.Users, error) {
	return ctrl.UsersService.GetUsers()
}

func (ctrl *UsersCtrl) GetUserById(id int) (*models.Users, error) {
	return ctrl.UsersService.GetUserById(id)
}

func (ctrl *UsersCtrl) CreateUser(user *models.Users) error {
	if ctrl.UsersService.CheckUser(user.U_login) {
		return ErrUserExists
	}
	id, err := ctrl.FavService.CreateFavourite(&models.Favourites{F_cntSerials: 0})
	if err != nil {
		return err
	}
	user.SetIdFavourites(id)
	err = ctrl.UsersService.CreateUser(user)
	if err != nil {
		ctrl.FavService.DeleteFavourite(id)
		return err
	}
	return nil
}

func (ctrl *UsersCtrl) UpdateUser(user *models.Users) error {
	return ctrl.UsersService.UpdateUser(user)
}

func (ctrl *UsersCtrl) DeleteUser(id int) error {
	user, _ := ctrl.GetUserById(id)

	err := ctrl.UsersService.DeleteUser(id)
	if err != nil {
		return err
	}

	return ctrl.FavService.DeleteFavourite(user.GetIdFavourites())
}

func (ctrl *UsersCtrl) GetUserByLogin(login string) (*models.Users, error) {
	return ctrl.UsersService.GetUserByLogin(login)
}

func (ctrl *UsersCtrl) AuthUser(login, password string) (*models.Users, error) {
	user, err := ctrl.UsersService.GetUserByLogin(login)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(user.U_password), []byte(password)) != nil {
		return nil, ErrInvalidPass
	}

	return user, nil
}

func (ctrl *UsersCtrl) GrantAdmin(id int) error {
	user, err := ctrl.UsersService.GetUserById(id)
	if err != nil {
		return ErrUserNotFound
	}

	user.U_role = "admin"

	return ctrl.UsersService.UpdateUser(user)
}
