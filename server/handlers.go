package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserId string `json:"id"`
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func login(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println("error when creating user")
		return
	}
	c.JSON(http.StatusOK, newUser.UserId)
}
