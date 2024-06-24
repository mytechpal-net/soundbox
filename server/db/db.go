package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	var err error
	// Init Db connection
	db, err = sql.Open("sqlite3", "./soundbox.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	log.Println("Connected to the database successfully!")

	CreateTables(db)
}

func Close() {
	if db != nil {
		db.Close()
	}
}

type UserProfile struct {
	Id     int64
	AuthId string
	Role   *string
	Token  *UserToken
}

type UserToken struct {
	UserId   string
	Token    string
	TokenExp time.Time
}

func GetUserProfile(authId string) *UserProfile {
	var user UserProfile

	err := db.QueryRow("SELECT id, authid, role from user where authid=$1", authId).Scan(&user.Id, &user.AuthId, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User nor found")
			return nil
		} else {
			log.Printf("%v", err)
			return nil
		}
	}

	return &user
}

func GetUserRole(authId string) string {
	var userRole string

	err := db.QueryRow("SELECT role from user where authid=$1", authId).Scan(&userRole)

	if err != nil {
		userRole = "user"
	}

	return userRole
}

/*
* Not really a good func : missing error management.
 */
func CreateUser(authId string) *UserProfile {
	query := "INSERT INTO user (authid, role) VALUES ($1, 'user')"
	result, err := db.Exec(query, authId)
	if err != nil {
		log.Println("unable to create user")
		log.Printf("error : %v", err)
	}

	id, _ := result.LastInsertId()

	log.Printf("User %v created, row id %v", authId, id)
	return &UserProfile{
		Id:     id,
		AuthId: authId,
	}
}

func PromoteUser(authId string) {
	query := "UPDATE user SET role = 'admin' WHERE authid = $1"
	_, err := db.Exec(query, authId)
	if err != nil {
		log.Println("unable to promote user")
		log.Printf("error : %v", err)
	}
}

func SaveToken(userId string, token UserToken) int64 {
	query := "INSERT INTO user_token (user_authid, token, token_exp) VALUES ($1, $2, $3) ON CONFLICT DO UPDATE SET token = $2, token_exp = $3"
	result, err := db.Exec(query, userId, token.Token, token.TokenExp)
	if err != nil {
		log.Println("Unable to save token")
		log.Printf("%v", err)
	}

	nbRows, _ := result.RowsAffected()
	return nbRows
}

func GetToken(token string) *UserToken {
	var userToken UserToken
	err := db.QueryRow("SELECT user_authid, token, token_exp from user_token WHERE token = $1", token).Scan(&userToken.UserId, &userToken.Token, &userToken.TokenExp)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return &userToken
}

func DelToken(token string) bool {
	_, err := db.Exec("DELETE FROM user_token WHERE token = $1", token)
	return err == nil
}

type SoundBox struct {
	Id        string
	Name      string
	Capacity  int
	Code      string
	SoundList []Sound
}

type Sound struct {
	Id         string
	Name       string
	SoundBoxId string
}

func GetSoundbox(id string) *SoundBox {
	var sb SoundBox
	err := db.QueryRow("SELECT id, name, code, capacity FROM soundbox WHERE id = $1", id).Scan(&sb.Id, &sb.Name, &sb.Code, &sb.Capacity)
	if err != nil {
		log.Println("Unable to get soundbox")
		return nil
	}

	sb.SoundList = GetSoundBoxSounds(id)

	return &sb
}

func CreateSoundBox(id string, name string, invitationCode string) error {
	query := "INSERT INTO soundbox (id, name, code, capacity) VALUES ($1, $2, $3, 50)"
	_, err := db.Exec(query, id, name, invitationCode)

	if err != nil {
		log.Println("Unable to create soundbox")
		log.Printf("%v", err)
		return err
	}

	return nil
}

func GetSoundboxByCode(code string) *SoundBox {
	var sb SoundBox
	err := db.QueryRow("SELECT id, name, code, capacity FROM soundbox WHERE code = $1", code).Scan(&sb.Id, &sb.Name, &sb.Code, &sb.Capacity)
	if err != nil {
		log.Println("Unable to get soundbox")
		return nil
	}

	return &sb
}

/*
Get the user sb
*/
func GetUserSb(userId string) *SoundBox {
	var sbId string
	err := db.QueryRow("SELECT soundbox_id FROM user_soundbox where user_authid = $1", userId).Scan(&sbId)
	if err != nil {
		log.Println("Looks like the user don't have any sb")
		return nil
	}

	return GetSoundbox(sbId)
}

func GetSoundBoxSounds(sbId string) []Sound {
	var sounds []Sound

	query := "SELECT id, name, soundbox_id FROM sound WHERE soundbox_id = $1"
	rows, err := db.Query(query, sbId)

	if err != nil {
		log.Println("Ooopsie")
	}

	for rows.Next() {
		var sound Sound
		rows.Scan(&sound.Id, &sound.Name, &sound.SoundBoxId)
		sounds = append(sounds, sound)
	}

	return sounds
}

/*
Join a sb with a code.
There is a Unique key on userId, then we need to check the result of the insert
*/
func JoinSoundBox(userId string, soundBoxCode string) *SoundBox {
	sb := GetSoundboxByCode(soundBoxCode)

	if sb == nil {
		log.Printf("No matching soundbox with the code %v\n", soundBoxCode)
		return nil
	}

	_, err := db.Exec("INSERT INTO user_soundbox (user_authid, soundbox_id) VALUES ($1, $2)", userId, sb.Id)
	if err != nil {
		log.Println("Unable to join the soundbox.")
		log.Printf("%v\n", err)
		return nil
	}

	sb.SoundList = GetSoundBoxSounds(sb.Id)

	return sb
}

func CreateSound(id string, name string, sbId string) error {

	_, err := db.Exec("INSERT INTO sound (id, name, soundbox_id) VALUES ($1, $2, $3)", id, name, sbId)
	if err != nil {
		log.Println("Unable to create sound entry")
		log.Printf("%v\n", err)
		return err
	}

	return nil
}
