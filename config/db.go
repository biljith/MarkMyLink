package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

//// database
//var DB *mgo.Database
//
//// collections
//var Bookmark *mgo.Collection

func init() {
	// get a mongo sessions
	// connecting to mongodb with authentication.
	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	//session, err := mgo.Dial("mongodb://localhost/BookMarkDB")
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	fmt.Println(databases)
	BookmarkDB := client.Database("BookmarkDB")
	Bookmark = BookmarkDB.Collection("Bookmarks")
}
