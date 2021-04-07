package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

// Manage the default API route (which is /)
func defaultRouters() {
	router.Any("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Unauthorized user"}) })
}

/* Default go function
initialize to all endpoints*/
func init() {
	defaultRouters()
	// Endpoints
	topSecret(router)
}

// Start: Run the gin router
func Start() {
	router.Run()
}
