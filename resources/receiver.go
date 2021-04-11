package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/juancsr/quasar-attack-operation/constants"
	"github.com/juancsr/quasar-attack-operation/controllers"
	"github.com/juancsr/quasar-attack-operation/models"
)

var body struct {
	Satellites []models.Satellite `json:"satellites" binding:"required"`
}

// common 404 error handler
var errorRespnse = func(c *gin.Context, r interface{}) { c.JSON(constants.NotFound_404, gin.H{"message": r}) }

// TopSecret endpoint handler
func TopSecret(c *gin.Context) {
	var response models.MotherShip

	if c.ShouldBind(&body) != nil {
		c.AbortWithStatus(constants.NotFound_404)
	}

	defer func() {
		if r := recover(); r != nil {
			errorRespnse(c, r)
		}
	}()

	decriptedInfo := controllers.DecriptMessage(body.Satellites)

	mothership := models.MotherShip{
		Position: models.Position{X: decriptedInfo["x"].(float32), Y: decriptedInfo["y"].(float32)},
		Message:  decriptedInfo["message"].(string),
	}

	response = mothership

	c.JSON(constants.OK_200, response)
}

func MultipleTopSecret(c *gin.Context) {

	if c.ShouldBind(&body) != nil {
		c.AbortWithStatus(constants.NotFound_404)
	}

	satellite_name := c.Param("satellite_name")

	defer func() {
		if r := recover(); r != nil {
			errorRespnse(c, r)
		}
	}()

	controllers.MultipleDecriptMessage(satellite_name)

	c.JSON(constants.OK_200, "ok")
}
