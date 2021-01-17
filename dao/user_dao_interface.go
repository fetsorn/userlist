package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDaoInterface interface {
	GetUser(filter interface{}) (result *mongo.SingleResult)
	GetUsers(filter interface{}) (cur *mongo.Cursor, err error)
}
