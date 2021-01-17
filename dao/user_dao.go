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

func (Collection *coll) GetUsers(filter interface{}) (cur *mongo.Cursor, err error) {
	cur, err = Collection.UserCollection.Find(context.TODO(), filter)
	return cur, err
}
