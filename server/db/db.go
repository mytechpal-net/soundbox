package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool
var dbErr error

func Init() {

	// Init Db connection
	dbPool, dbErr = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", dbErr)
		os.Exit(1)
	}

	log.Println("Connected to the database successfully!")
}

func Close() {
	if dbPool != nil {
		dbPool.Close()
	}
}

type UserProfile struct {
	Id       int
	AuthId   string
	Nickname *string
	Role     []*string
	Token    *UserToken
}

type UserToken struct {
	UserId   string
	Token    string
	TokenExp time.Time
}

func GetUserProfile(authId string) *UserProfile {
	var user UserProfile

	err := dbPool.QueryRow(context.Background(), "select id, authid, nickname, role from \"users\" where authid=$1", authId).Scan(&user.Id, &user.AuthId, &user.Nickname, &user.Role)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return &user
}

/*
* Not really a good func : missing error management.
 */
func CreateUser(authId string) *UserProfile {
	var id int
	query := "INSERT INTO \"users\" (authid, role) VALUES ($1, '{''user''}') returning id"
	err := dbPool.QueryRow(context.Background(), query, authId).Scan(&id)
	if err != nil {
		log.Println("unable to create user")
		log.Printf("error : %v", err)
	}

	return &UserProfile{
		Id:     id,
		AuthId: authId,
	}
}

func SaveToken(userId int, token UserToken) int {
	var id int
	query := "INSERT INTO \"users_token\" (user_id, token, token_exp) values ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET token = $2, token_exp = $3 returning user_id"
	err := dbPool.QueryRow(context.Background(), query, userId, token.Token, token.TokenExp).Scan(&id)
	if err != nil {
		log.Println("Unable to save token")
		log.Printf("%v", err)
	}

	return id
}

func GetToken(token string) *UserToken {
	var userToken UserToken
	err := dbPool.QueryRow(context.Background(), "select user_id, token, token_exp from \"users_token\" where token = $1", token).Scan(&userToken.UserId, &userToken.Token, &userToken.TokenExp)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	return &userToken
}

type soundBox struct {
	Id       int
	Name     string
	Capacity int
	Code     string
}

func GetSoundbox(id int) *soundBox {
	var sb soundBox
	err := dbPool.QueryRow(context.Background(), "SELECT id, name, code, capacity FROM soundbox WHERE id = $1", id).Scan(&sb.Id, &sb.Name, &sb.Code, &sb.Capacity)
	if err != nil {
		log.Println("Unable to get soundbox")
		return nil
	}

	return &sb
}

func GetSoundboxByCode(code string) *soundBox {
	var sb soundBox
	err := dbPool.QueryRow(context.Background(), "SELECT id, name, code, capacity FROM soundbox WHERE code = $1", code).Scan(&sb.Id, &sb.Name, &sb.Code, &sb.Capacity)
	if err != nil {
		log.Println("Unable to get soundbox")
		return nil
	}

	return &sb
}

/*
Get the user sb
*/
func GetUserSb(userId string) *soundBox {
	var sbId int
	err := dbPool.QueryRow(context.Background(), "SELECT soundbox_id FROM user_soundbox where user_authid = $1", userId).Scan(&sbId)
	if err != nil {
		log.Println("Looks like the user don't have any sb")
		return nil
	}

	return GetSoundbox(sbId)
}

/*
Join a sb with a code.
There is a Unique key on userId, then we need to check the result of the insert
*/
func JoinSoundBox(userId string, soundBoxCode string) bool {
	sb := GetSoundboxByCode(soundBoxCode)

	if sb == nil {
		log.Printf("No matchin soundbox with the code %v\n", soundBoxCode)
		return false
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO user_soundbox (user_authid, soundbox_id) VALUES ($1, $2) returning soundbox_id", userId, sb.Id)
	if err != nil {
		log.Println("Unable to join the soundbox.")
		log.Printf("%v\n", err)
		return false
	}

	return true
}
