package models

import (
	"errors"
	"log"

	"github.com/juancsr/quasar-attack-operation/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SatelliteRequest struct {
	Name     string   `bson:"name" json:"name" binding:"required"`
	Distance float32  `bson:"distance" json:"distance" binding:"required"`
	Message  []string `bson:"message" json:"message" binding:"required"`
}

const collectionName = "requests"

// search for only one request using the satellite name as key filter
func GetOneRequestBySatelliteName(newRequest SatelliteRequest) (*SatelliteRequest, error) {
	var request SatelliteRequest

	ctx := db.CTX
	collection := db.GetCollection(collectionName)
	filter := bson.M{"name": newRequest.Name}
	update := bson.M{
		"$set": bson.M{
			"name":     newRequest.Name,
			"distance": newRequest.Distance,
			"message":  newRequest.Message,
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&request)
	//defer db.CloseConnection()
	return &request, err
}

// Retrive all the requests
func GetAllSatelliteRequest() []SatelliteRequest {
	var results []SatelliteRequest
	ctx := db.CTX
	cursor, err := db.GetCollection(collectionName).Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		log.Panic(errors.New("not enough information"))
	}
	err = cursor.All(ctx, &results)
	if err != nil {
		log.Println(err)
		log.Panic(errors.New("not enough information"))
	}
	return results
}
