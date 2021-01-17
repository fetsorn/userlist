package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"userlist/models"
)

type UserService interface {
	DeleteUserService(id primitive.ObjectID) (res *mongo.DeleteResult, err error)
	GetUserService(id primitive.ObjectID) (user models.User, err error)
	GetUsersService() (Users []models.User, err error)
}
