package controllers

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPass  = errors.New("invalid password")
	ErrUserExists   = errors.New("user already exists")
)
