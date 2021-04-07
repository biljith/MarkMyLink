package model



import (
	"MarkMyLink/config"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	//"log"
	//"fmt"
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

	//if err = bmCursor.All(context.TODO(), &bm); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Hi")
	//fmt.Println(bm)
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
