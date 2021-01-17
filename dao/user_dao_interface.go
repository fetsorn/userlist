package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"userlist/models"
)

type UserDaoInterface interface {
	AddUser(u *models.User) (result *mongo.InsertOneResult, err error)
	DeleteUser(filter interface{}) (result *mongo.DeleteResult, err error)
	GetUser(filter interface{}) (result *mongo.SingleResult)
	GetUsers(filter interface{}) (cur *mongo.Cursor, err error)
}
