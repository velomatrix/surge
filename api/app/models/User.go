package models

import (
	"fmt"
	"time"
	"math/rand"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	AuthToken string	`json:"auth_token" bson:"auth_token"`
	CreatedAt time.Time	`json:"created_at" bson:"creatd_at"`
	UpdatedAt time.Time	`json:"updated_at" bson:"updated_at"`
}

const (
	USERS_COLLECTION = "users"
	AUTH_TOKEN_ALPHABET = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	AUTH_TOKEN_LENGTH = 40
)

func FindUser(s *mgo.Session, id bson.ObjectId) *User {
	var user User
	s.DB("").C(USERS_COLLECTION).FindId(id).One(&user)
	return &user
}

func FindAllUsers(s *mgo.Session) (*[]User, error) {
	var results []User

	err := s.DB("").C(USERS_COLLECTION).Find(nil).All(&results)
	if err != nil {
		fmt.Print("Unable to retrieve all Ausers")
		return nil, err
	}
	return &results, nil
}

func NewUser(s *mgo.Session, username string) (*User, error) {
	now := time.Now()
	user := User{
		Username: username,
		AuthToken: genAuthToken(AUTH_TOKEN_LENGTH),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.DB("").C(USERS_COLLECTION).Insert(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func genAuthToken(tokenLen int) string {
	buf := make([]byte, tokenLen)
	for i := 0; i < tokenLen; i++ {
		buf[i] = AUTH_TOKEN_ALPHABET[rand.Intn(len(AUTH_TOKEN_ALPHABET))]
	}

	return string(buf)
}
