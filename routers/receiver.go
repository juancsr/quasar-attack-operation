package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/juancsr/quasar-attack-operation/resources"
)

// add all of the /topsecret routers
func topSecret(router *gin.Engine) {
	router.GET("/topsecret/:satellite_name", resources.MultipleTopSecret)
	router.POST("/topsecret", resources.TopSecret)
	router.POST("/topsecret/:satellite_name", resources.MultipleTopSecret)
}
