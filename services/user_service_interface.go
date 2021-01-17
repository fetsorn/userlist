package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"userlist/models"
)

type UserService interface {
	GetUserService(id primitive.ObjectID) (user models.User, err error)
	GetUsersService() (Users []models.User, err error)
}
