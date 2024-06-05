package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmathelot/soundbox-server/db"
	"google.golang.org/api/idtoken"
)

type user struct {
	ClientId   string `json:"clientId"`
	Credential string `json:"credential"`
}

type userData struct {
	UserId     string
	SessionExp int
}

func login(c *gin.Context) {
	var userCredentials user

	// bind post values
	if err := c.BindJSON(&userCredentials); err != nil {
		fmt.Printf("Error %v", err)
		c.JSON(http.StatusBadRequest, "Incorrect login values")
		return
	}

	ctx := context.Background()

	// Validate the token
	payload, err := validateGoogleToken(ctx, userCredentials.ClientId, userCredentials.Credential)
	if err != nil {
		c.JSON(http.StatusForbidden, "Token validation failed")
		return
	}

	user := db.GetUserProfile(payload.Subject)
	if user == nil {
		user = db.CreateUser(payload.Subject)
	}

	user.Token = generateToken(user.AuthId)

	log.Println(user.Token)

	// Save token & write cookie
	db.SaveToken(user.Id, *user.Token)
	c.SetCookie("sb_session", user.Token.Token, 28800, "/", "localhost", true, true)

	c.JSON(http.StatusOK, userData{
		user.AuthId,
		int(user.Token.TokenExp.UnixMilli()),
	})
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// validateGoogleToken validates the given Google ID token
func validateGoogleToken(ctx context.Context, audience string, token string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, token, audience)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// generate a session token
func generateToken(userId string) *db.UserToken {
	b := make([]byte, 8)
	rand.Read(b)

	t0 := time.Now().Add(time.Hour * 8)

	return &db.UserToken{
		UserId:   userId,
		Token:    fmt.Sprintf("%x", b),
		TokenExp: t0,
	}
}
