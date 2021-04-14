package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/juancsr/quasar-attack-operation/constants"
	"github.com/juancsr/quasar-attack-operation/controllers"
	"github.com/juancsr/quasar-attack-operation/models"
)

// common 404 error handler
var errorResponse = func(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(constants.NotFound_404, gin.H{"message": r})
	}
}

// handle the common response of /topsecret and /topsecret/{satellite_name}
// by assign the values of decriptedInfo map into a Mothership struct
func decriptedInfoResponseHandler(decriptedInfo map[string]interface{}) (mothership models.MotherShip) {
	mothership = models.MotherShip{
		Position: models.Position{X: decriptedInfo["x"].(float32), Y: decriptedInfo["y"].(float32)},
		Message:  decriptedInfo["message"].(string),
	}
	return
}

// TopSecret godoc
// @Summary Top secret endpoint
// @Description trilateration point
// @Accept  json
// @Produce  json
// @Param satellite_name body models.SatelliteRequest true "Satelline name"
// @Success 200 {object} models.MotherShip
// @Failure 400,404 {object}  models.MotherShip
// @Failure default {object}  models.MotherShip
// @Router /topsecret [post]
func TopSecret(c *gin.Context) {
	var body struct {
		Satellites []models.SatelliteRequest `json:"satellites" binding:"required"`
	}

	if c.ShouldBind(&body) != nil {
		c.AbortWithStatus(constants.NotFound_404)
	}

	defer errorResponse(c)

	decriptedInfo := controllers.DecriptMessage(body.Satellites)
	response := decriptedInfoResponseHandler(decriptedInfo)

	c.JSON(constants.OK_200, response)
}

// MultipleTopSecret godoc
// @Summary MultipleTopSecret
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param satellite_name path int true "Satelline name"
// @Success 200 {object} models.MotherShip
// @Failure 400,404 {object}  models.MotherShip
// @Failure default {object}  models.MotherShip
// @Router /topscret/{satellite_name} [post][get]
func MultipleTopSecret(c *gin.Context) {
	var satelliteRequests models.SatelliteRequest

	defer errorResponse(c)

	satellite_name := c.Param("satellite_name")

	if c.Request.Method == "POST" {
		var body struct {
			Distance float32  `json:"distance" binding:"required"`
			Message  []string `json:"message" binding:"required"`
		}

		if c.ShouldBind(&body) != nil {
			c.AbortWithStatus(constants.NotFound_404)
		}
		satelliteRequests = models.SatelliteRequest{
			Name:     satellite_name,
			Distance: body.Distance,
			Message:  body.Message,
		}
	}

	decriptedInfo := controllers.MultipleDecriptMessage(satelliteRequests)
	response := decriptedInfoResponseHandler(decriptedInfo)

	c.JSON(constants.OK_200, response)
}
