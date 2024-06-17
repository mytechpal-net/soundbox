package main

import (
	"github.com/gin-contrib/cors"
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
	app.GET("/user/logout", logout)
	app.POST("/user/join", joinSoundBox)

	app.GET("/soundbox/:id", soundBox)

	router.Run()
}

// Create gin router
func newRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://soundbox.mytechpal.net"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	return router
}
