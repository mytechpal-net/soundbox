package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	router := newRouter()

	// add routes
	router.GET("/ping", ping)

	router.Run()
}

// Create gin router
func newRouter() *gin.Engine {
	router := gin.Default()

	trustedProxies := []string{"127.0.0.1"}

	router.SetTrustedProxies(trustedProxies)

	return router
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
