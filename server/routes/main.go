package routes

import (
	"github.com/gin-gonic/gin"
)

var e = gin.Default()

// Run will start the server
func Run() {
	// add routes handler
	addPingRoutes(e)

	e.Run(":5000")
}
