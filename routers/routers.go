package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

func defaultRouters() {
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Unauthorized user"}) })
	router.POST("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Unauthorized user"}) })
}

/* Default go function
initialize to all endpoints*/
func init() {
	// Default routes
	defaultRouters()
	// Endpoints
	topSecret(router)
}

// Start: Run the gin router
func Start() {
	router.Run()
}
