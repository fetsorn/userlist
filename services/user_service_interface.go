package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"userlist/models"
)

type UserService interface {
	AddUserService(u *models.User) (res *mongo.InsertOneResult, err error)
	DeleteUserService(id primitive.ObjectID) (res *mongo.DeleteResult, err error)
	GetUserService(id primitive.ObjectID) (user models.User, err error)
	GetUserByFirstService(p string) (user models.User, err error)
	GetUserByLastService(p string) (user models.User, err error)
	GetUserByCityService(p string) (user models.User, err error)
	GetUserByCountryService(p string) (user models.User, err error)
	GetUsersService() (Users []models.User, err error)
}
