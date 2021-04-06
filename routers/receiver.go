package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/juancsr/quasar-attack-operation/controllers"
)

func topSecret(router *gin.Engine) {
	router.POST("/topsecret", controllers.DecriptMessage)
}
