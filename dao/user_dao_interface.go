package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDaoInterface interface {
	GetUsers(filter interface{}) (cur *mongo.Cursor, err error)
}
