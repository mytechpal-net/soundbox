package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
