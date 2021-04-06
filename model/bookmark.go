package model



import (
<<<<<<< HEAD
"MarkMyLink/config"
"go.mongodb.org/mongo-driver/bson"
//"go.mongodb.org/mongo-driver/mongo"
"context"
"log"
"fmt"
//"gopkg.in/mgo.v2/bson"
=======
	"MarkMyLink/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
>>>>>>> 87741d62c470c48405d5999fc6f2016be38939da
)

// BookMark struct()
type Bookmark struct {
	Name       string        `bson:"Name"`
	Link       string        `bson:"Link"`
	Viewcount  string		 `bson:"Viewcount"`
	Timestamp  string        `bson:"Timestamp"`
}

func AllBookmarks() ([]Bookmark, error) {

	bm := []Bookmark{}

	if err = bmCursor.All(context.TODO(), &bm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hi")
	fmt.Println(bm)
	return bm, nil
}

func FindBookmarks(email string) ([]Bookmark, error) {
	var bookmarks []Bookmark
	cursor, err := config.BookmarkCollection.Find(context.TODO(),
										 bson.D{{"email", email}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var bookmark Bookmark
		if err := cursor.Decode(&bookmark); err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}
