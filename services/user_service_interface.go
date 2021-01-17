package services

import (
	"userlist/models"
)

type UserService interface {
	GetUsersService() (Users []models.User, err error)
}
