package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDaoInterface interface {
	DeleteUser(filter interface{}) (result *mongo.DeleteResult, err error)
	GetUser(filter interface{}) (result *mongo.SingleResult)
	GetUsers(filter interface{}) (cur *mongo.Cursor, err error)
}
