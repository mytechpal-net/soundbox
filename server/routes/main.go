package routes

import (
	"github.com/gin-gonic/gin"
)

// Add routes to the main engine
func AddRoutes(e *gin.Engine) {
	addPingRoutes(e)
}
