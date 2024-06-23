package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmathelot/soundbox-server/db"
)

// Check token validity
func validateAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("sb_session")
		if err != nil {
			c.JSON(http.StatusForbidden, "no session found")
			c.Abort()
		}

		// Check token validity
		log.Println("Checking validity", cookie)
		token := db.GetToken(cookie)

		if token == nil || (token.TokenExp.Compare(time.Now()) == -1) {
			log.Println("Oh no")
			c.JSON(http.StatusForbidden, "no session found")
			c.Abort()
		}

		// token session is still valid we should potentially increase the validity
		c.Next()
	}

}

func canPlayTheFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
