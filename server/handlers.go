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

// Get soundbox per user
// SB can be nil on purpose.
func userContext(c *gin.Context) {
	userSb := db.GetUserSb(c.Param("authid"))
	c.JSON(http.StatusOK, userSb)
}

func soundBox(c *gin.Context) {
	log.Println(c.Param("id"))
	c.JSON(http.StatusOK, "OK")
}

type joinSbData struct {
	SoundBoxCode string `json:"invitationCode"`
	UserId       string `json:"user"`
}

func joinSoundBox(c *gin.Context) {
	var joinData joinSbData

	// bind post values
	if err := c.BindJSON(&joinData); err != nil {
		fmt.Printf("Error %v", err)
		c.JSON(http.StatusBadRequest, "Incorrect values")
		return
	}

	sb := db.JoinSoundBox(joinData.UserId, joinData.SoundBoxCode)
	if sb == nil {
		c.JSON(http.StatusBadRequest, "Incorrect values")
		return
	}

	c.JSON(http.StatusOK, sb)
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
