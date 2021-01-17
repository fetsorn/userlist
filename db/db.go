package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"userlist/models"
)

var (
	MongoDb *mongo.Database
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost",
	))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDb")
	MongoDb = client.Database("userlist")

	// mock one entry for development
	coll := MongoDb.Collection("Users")
	coll.DeleteMany(context.TODO(), bson.D{})

	u := models.User{First: "Sam", Last: "Jackson", Age: 34, City: "New York", Country: "USA"}
	_, err = coll.InsertOne(context.TODO(), u)
	if err != nil {
		return
	}
}

func GetUserCollection() *mongo.Collection {
	return MongoDb.Collection("Users")
}
