package model



import (
	"MarkMyLink/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// BookMark struct()
type Bookmark struct {
	Name       string        `bson:"Name"`
	Link       string        `bson:"Link"`
	Viewcount 	int			 `bson:"Viewcount"`
	Timestamp  string        `bson:"Timestamp"`
}

func AllBookmarks() ([]Bookmark, error) {

	bm := []Bookmark{}

	//err := config.Bookmark.Find(bson.M{}).All(&bm)
	//if err != nil {
	//	return nil, err
	//}
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
