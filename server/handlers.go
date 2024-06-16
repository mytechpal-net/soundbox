package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
		log.Printf("User %v not found, creating it", payload.Subject)
		user = db.CreateUser(payload.Subject)
	}

	user.Token = generateToken(user.AuthId)

	// Save token & write cookie
	db.SaveToken(user.Id, *user.Token)
	c.SetCookie("sb_session", user.Token.Token, 28800, "/", "localhost", true, true)

	c.JSON(http.StatusOK, userData{
		user.AuthId,
		int(user.Token.TokenExp.UnixMilli()),
	})
}

func logout(c *gin.Context) {
	cookie, _ := c.Cookie("sb_session")
	cleaned := db.DelToken(cookie)
	c.SetCookie("sb_session", cookie, 0, "/", "localhost", true, true)
	c.JSON(http.StatusOK, cleaned)
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connection structure
type Connection struct {
	Conn  *websocket.Conn
	Group string
}

var connections = make(map[*websocket.Conn]*Connection)

func soundBox(c *gin.Context) {

	group := c.Param("id")
	if group == "" {
		c.JSON(http.StatusPreconditionFailed, "Group not specified")
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer conn.Close()

	connection := &Connection{Conn: conn, Group: group}
	connections[conn] = connection

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			delete(connections, conn)
			break
		}

		log.Printf("received from %s: %s", group, message)

		// Broadcast message to other clients in the same group
		for _, c := range connections {
			if c.Group == group && c.Conn != conn {
				log.Println("sending message:", message)
				err = c.Conn.WriteMessage(messageType, message)
				if err != nil {
					log.Println("write:", err)
					c.Conn.Close()
					delete(connections, c.Conn)
				}
			}
		}
	}
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
