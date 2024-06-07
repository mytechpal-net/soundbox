package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kmathelot/soundbox-server/db"
)

func main() {
	// init Db connexion
	db.Init()
	defer db.Close()

	// Initialize router
	router := newRouter()

	router.POST("/login", login)

	// Put the app routes under a group
	app := router.Group("/app")
	// Check session validity
	app.Use(validateAuthorizationMiddleware())

	app.GET("/ping", pong)
	app.GET("/user/:authid", userContext)

	app.GET("/soundbox/:id", soundBox)
	app.POST("/sounbox/join", joinSoundBox)

	router.Run()
}

// Create gin router
func newRouter() *gin.Engine {
	router := gin.Default()

	trustedProxies := []string{"127.0.0.1"}

	router.SetTrustedProxies(trustedProxies)

	return router
}
