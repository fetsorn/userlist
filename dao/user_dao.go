package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"userlist/db"
)

type coll struct {
	UserCollection *mongo.Collection
}

var Collection *coll

func NewDao() *coll {
	Collection = &coll{UserCollection: db.GetUserCollection()}
	return Collection
}

func (Collection *coll) DeleteUser(filter interface{}) (result *mongo.DeleteResult, err error) {
	result, err = Collection.UserCollection.DeleteOne(context.TODO(), filter)
	return result, err
}

func (Collection *coll) GetUser(filter interface{}) (result *mongo.SingleResult) {
	result = Collection.UserCollection.FindOne(context.TODO(), filter)
	return result
}

func (Collection *coll) GetUsers(filter interface{}) (cur *mongo.Cursor, err error) {
	cur, err = Collection.UserCollection.Find(context.TODO(), filter)
	return cur, err
}
