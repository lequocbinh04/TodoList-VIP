package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions *options.ClientOptions
var client *mongo.Client
var err error

func InitDatabase() {
	clientOptions = options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
    	log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
    	log.Fatal(err)
	}
}

func ConnectCollection(nameCollection string) *mongo.Collection {
	collection := client.Database("todolist").Collection(nameCollection)
	return collection
}
