package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/juancsr/quasar-attack-operation/resources"
)

func topSecret(router *gin.Engine) {
	router.POST("/topsecret", resources.TopSecret)
	router.POST("/topsecret/:satellite_name", resources.MultipleTopSecret)
}
