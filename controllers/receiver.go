package controllers

import (
	"errors"
	"log"

	"github.com/juancsr/quasar-attack-operation/models"
	"github.com/juancsr/quasar-attack-operation/utils"
)

// Decript the message using the satellites and the message
func DecriptMessage(satellites []models.SatelliteRequest) map[string]interface{} {
	var decriptedInfo = make(map[string]interface{})

	var distances = []float32{}
	var messages = [][]string{}

	for _, satellite := range satellites {
		distances = append(distances, satellite.Distance)
		messages = append(messages, satellite.Message)
	}

	x, y := utils.GetLocation(distances...)
	decriptedInfo["x"] = x
	decriptedInfo["y"] = y

	decriptedInfo["message"] = utils.GetMessage(messages...)
	return decriptedInfo
}

// Check if the request already exist in the db.
// If not exist, then it will create it
// otherwise the data will be updated
func MultipleDecriptMessage(updatedRequest models.SatelliteRequest) map[string]interface{} {
	if updatedRequest.Name != "" {
		_, err := models.GetOneRequestBySatelliteName(updatedRequest)
		if err != nil {
			log.Println(err)
			log.Panic(errors.New("not enough information"))
		}
	}

	requests := models.GetAllSatelliteRequest()
	decriptedInfo := DecriptMessage(requests)
	return decriptedInfo
}
