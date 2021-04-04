package model

import (
	"MarkMyLink/config"

	"go.mongodb.org/mongo-driver/bson"
	"context"
	"log"
)

// BookMark struct()
type User struct {
	Email      string        `bson:"email"`
	Password   string        `bson:"password"`
	Name       string        `bson:"name"`
}


// Create a new user in the collection Users. Returns false if user with same email
// already exists. True otherwise.
func CreateUser(user *User) bool {
	log.Printf("Creating user")

	var userWithSameEmail User
	err := config.UserCollection.FindOne(context.TODO(),
										 bson.D{{"email", user.Email}}).Decode(&userWithSameEmail)
	if err == nil {
		// A user with same email already exists.
		return false
	} else {
		// No user with same email. Create one.
		if _, err := config.UserCollection.InsertOne(context.TODO(), user); err != nil {
		 	return false
		} else {
			return true
		}
	}
}

func FindUser(email string) (User, error) {
	var user User
	err := config.UserCollection.FindOne(context.TODO(),
										 bson.D{{"email", email}}).Decode(&user)
	return user, err
}
