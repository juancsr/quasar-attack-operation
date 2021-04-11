package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CTX Context variable
var CTX = context.TODO()

var client *mongo.Client

// var collection *mongo.Collection

// init connection
func init() {
	var err error
	// Docker URI
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	DB := os.Getenv("DB")
	CLUSTER := os.Getenv("CLUSTER")
	URI := "mongodb+srv://" + USER + ":" + PASSWORD + "@" + CLUSTER + "/" + DB + "?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(URI)
	client, err = mongo.Connect(CTX, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(CTX, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// cliente de db
func GetCollection(collectionName string) *mongo.Collection {
	client.Connect(CTX)
	collection := client.Database(os.Getenv("DB")).Collection(collectionName)
	return collection
}

// cierra la conexión con la BD
func CloseConnection() {
	defer func() {
		if err := client.Disconnect(CTX); err != nil {
			log.Panicln(err.Error())
		}
	}()
}
