package database

import (
	"context"
	"log"
	"mongo_go/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection
var ctx = context.TODO()

const db = "training_db"

var collection = "users"

func createConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	client := createConnection()
	collection := client.Database(db).Collection(collection)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func FindUser(username string) (models.User, error) {
	client := createConnection()
	collection := client.Database(db).Collection(collection)

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	var result models.User

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func UpdateUser(user models.User) (*mongo.UpdateResult, error) {
	client := createConnection()
	collection := client.Database(db).Collection(collection)

	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, user)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeleteUser(username string) (*mongo.DeleteResult, error) {
	client := createConnection()
	collection := client.Database(db).Collection(collection)

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	deleteResult, err := collection.DeleteMany(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
