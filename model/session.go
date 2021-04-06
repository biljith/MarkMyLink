package model

import (
	"MarkMyLink/config"

	"time"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Session struct {
	CreatedAt time.Time `bson:"createdAt"`
	Email string        `bson:"email"`
	Token string        `bson:token`
}

func CreateSession(session *Session) bool {
	log.Printf("Creating session")
	// No user with same email. Create one.
	if _, err := config.SessionCollection.InsertOne(context.TODO(), session); err != nil {
	 	return false
	} else {
		return true
	}
}

func FindSession(sessionToken string) (Session, error) {
	var session Session
	err := config.SessionCollection.FindOne(context.TODO(),
										 bson.D{{"token", sessionToken}}).Decode(&session)
	return session, err
}