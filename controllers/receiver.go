package controllers

import (
	"fmt"

	"github.com/juancsr/quasar-attack-operation/models"
	"github.com/juancsr/quasar-attack-operation/utils"
)

// Decript the message using the satellites and the message
func DecriptMessage(satellites []models.Satellite) map[string]interface{} {
	var decriptedInfo = make(map[string]interface{})

	var distances = []float32{}
	var messages = [][]string{}

	for _, satellite := range satellites {
		distances = append(distances, satellite.Distance)
		messages = append(messages, satellite.Message)
	}

	//x, y := utils.GetLocation(distances...)
	decriptedInfo["x"] = float32(1.0)
	decriptedInfo["y"] = float32(2.0)

	decriptedInfo["message"] = utils.GetMessage(messages...)

	return decriptedInfo
}

// Check if the request already exist in the db.
// If not exist, then it will create it
// otherwise the data will be updated
func MultipleDecriptMessage(satellite_name string) {
	fmt.Println(satellite_name)
}
