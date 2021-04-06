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

var BookmarkCollection *mongo.Collection
var CTX context.Context
var UserCollection *mongo.Collection
var SessionCollection *mongo.Collection
var BookmarkCollection *mongo.Collection

func init() {
	// get a mongo sessions
	// connecting to mongodb with authentication.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://diddlz:terimarzi69@cluster0.l0c9p.mongodb.net/BookmarkDB?retryWrites=true&w=majority"))
	//session, err := mgo.Dial("mongodb://localhost/BookMarkDB")
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	CTX = ctx
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
	// BookmarkDB := client.Database("BookmarkDB")
	BookmarkCollection = client.Database("BookmarkDB").Collection("Bookmarks")
	UserCollection = client.Database("BookmarkDB").Collection("Users")
	SessionCollection = client.Database("BookmarkDB").Collection("Sessions")
	fmt.Println(BookmarkCollection)
	fmt.Println(UserCollection)
	fmt.Println(SessionCollection)

	// Writing to DB

	// BookmarkResult, err := BookmarkCollection.InsertOne(ctx, bson.D{
	// 	{Key: "Name" , Value: "Twitter"},
	// 	{Key: "Link" , Value: "www.twitter.com"},
	// 	{Key: "ViewCount" , Value: 1},
	// 	{Key: "Timestamp" , Value:"2020-11-10 23:00:00 +0000 UTC m=+0.000000000"},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(BookmarkResult.InsertedID)

	// // Reading from DB

	// cursor, err := BookmarkCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	var bookmark bson.M
	// 	if err = cursor.Decode(&bookmark); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(bookmark)
	// }

	// // Filter Docs with attributes
	// filterCursor, err := BookmarkCollection.Find(ctx, bson.M{"Name": "twitter"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var bookmarksFiltered []bson.M
	// if err = filterCursor.All(ctx, &bookmarksFiltered); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(bookmarksFiltered)
}
