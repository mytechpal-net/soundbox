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

func Init() {
	var err error
	// Init Db connection
	dbPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
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
