package controllers

import (
	codes "github.com/juancsr/quasar-attack-operation/constants"

	"github.com/gin-gonic/gin"
)

func DecriptMessage(c *gin.Context) {
	c.JSON(codes.OK, gin.H{"message": "OK"})
}
