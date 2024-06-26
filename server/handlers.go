package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/kmathelot/soundbox-server/db"
	"github.com/kmathelot/soundbox-server/internal/directories"
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

	user.Token = generateSessionToken(user.AuthId)

	// Save token & write cookie
	db.SaveToken(user.AuthId, *user.Token)
	c.SetCookie("sb_session", user.Token.Token, 28800, "/", c.Request.Host, true, true)

	c.JSON(http.StatusOK, userData{
		user.AuthId,
		int(user.Token.TokenExp.UnixMilli()),
	})
}

func logout(c *gin.Context) {
	cookie, _ := c.Cookie("sb_session")
	cleaned := db.DelToken(cookie)
	c.SetCookie("sb_session", cookie, 0, "/", c.Request.Host, true, true)
	c.JSON(http.StatusOK, cleaned)
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

type userSbContext struct {
	SoundBox *db.SoundBox
	UserRole string
}

// Get soundbox per user
// SB can be nil on purpose.
// Return user Role (admin|user)
// BY default nobody is admin
func userContext(c *gin.Context) {
	userSb := db.GetUserSb(c.Param("authid"))

	userRole := "user"

	if userSb != nil {
		userRole = db.GetUserRole(c.Param("authid"))
	}

	userContext := userSbContext{
		userSb,
		userRole,
	}

	c.JSON(http.StatusOK, userContext)
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

var stats = make(map[string]int32)

func soundBox(c *gin.Context) {

	group := c.Param("id")
	if group == "" {
		c.JSON(http.StatusPreconditionFailed, "Group not specified")
		return
	}

	// Check if user can join the group
	cookie, _ := c.Cookie("sb_session")
	userToken := db.GetToken(cookie)
	if !db.IsUserAuthorized(userToken.UserId, group) {
		c.String(http.StatusPreconditionFailed, "You can't do that")
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
		stats[group] += 1

		// Broadcast message to other clients in the same group
		for _, w := range connections {
			if w.Group == group {
				log.Println("sending message:", message)
				err = w.Conn.WriteMessage(messageType, message)
				if err != nil {
					log.Println("write:", err)
					w.Conn.Close()
					delete(connections, w.Conn)
				}
			}
		}
	}
}

type newSbData struct {
	SbName string `json:"sbName"`
	UserId string `json:"user"`
}

func createSoundBox(c *gin.Context) {
	var sbData newSbData

	// bind post values
	if err := c.BindJSON(&sbData); err != nil {
		fmt.Printf("Error %v", err)
		c.JSON(http.StatusBadRequest, "Incorrect values")
		return
	}

	// Check if the user is not a member already
	if check := db.GetUserSb(sbData.UserId); check != nil {
		c.JSON(http.StatusForbidden, "msg: 'Hell no'")
		return
	}

	// If no name was provided, we generate a sequence
	if len(strings.TrimSpace(sbData.SbName)) == 0 {
		sbData.SbName = fmt.Sprintf("sb-%v", generateId(8))
	}

	sbId := generateId(8)
	invitationCode := fmt.Sprintf("%v-%v", generateId(3), generateId(3))

	// Create the SB
	if err := db.CreateSoundBox(sbId, sbData.SbName, invitationCode); err != nil {
		fmt.Printf("Error %v", err)
		c.JSON(http.StatusInternalServerError, "Unable to create the box")
		return
	}

	// Join the soundbox
	sb := db.JoinSoundBox(sbData.UserId, invitationCode)

	// Create the directory
	directories.CreateDirectory(fmt.Sprintf("./sounds/%v", sbId))

	// promote the user
	db.PromoteUser(sbData.UserId)

	// Return the new user context
	userSbCtx := userSbContext{
		sb,
		"admin",
	}

	c.JSON(http.StatusOK, userSbCtx)
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
func generateSessionToken(userId string) *db.UserToken {
	t0 := time.Now().Add(time.Hour * 8)

	return &db.UserToken{
		UserId:   userId,
		Token:    generateId(10),
		TokenExp: t0,
	}
}

// Create a rand Id
func generateId(length int) string {
	b := make([]byte, length)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

func uploadFile(c *gin.Context) {

	// Get user sb id
	// The access has already been granted by the middleware
	cookie, _ := c.Cookie("sb_session")
	userToken := db.GetToken(cookie)

	// Get user SB
	userSb := db.GetUserSb(userToken.UserId)

	// This case shouldn't happen
	if userSb == nil {
		c.String(http.StatusBadRequest, "Unable to process")
		return
	}

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Get form err: %s", err.Error())
		return
	}

	// Limit file size (e.g., 10MB)
	const maxFileSize = 10 << 20 // 10MB
	if file.Size > maxFileSize {
		c.String(http.StatusBadRequest, "File too large")
		return
	}

	name := c.PostForm("name")

	if len(name) == 0 {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// Validate file extension
	validExtensions := map[string]bool{
		".mp3": true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !validExtensions[ext] {
		c.String(http.StatusBadRequest, "Invalid file extension")
		return
	}

	// Create a unique file name
	uniqueFilename := uuid.New().String() + ext

	// Upload the file to specific dst.
	// Should resample de file to have a reduced size
	if err := c.SaveUploadedFile(file, fmt.Sprintf("./sounds/%s/%s", userSb.Id, uniqueFilename)); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}

	// Create an entry in the db
	if err := db.CreateSound(uniqueFilename, name, userSb.Id); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}

	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
