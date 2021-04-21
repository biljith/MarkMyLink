package model

import (
	"MarkMyLink/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	//"log"
	//"fmt"
)

// BookMark struct()
type Bookmark struct {
	Email  string		 `bson:"email"`
	Name       string        `bson:"name"`
	Link       string        `bson:"link"`
	Viewcount  int		 `bson:"viewcount"`
	Timestamp  string        `bson:"timestamp"`
	Image string `bson:"image"`
	Description string `bson:"description"`
	Category string `bson:"category"`
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

// Create a new bookmark in the collection Bookmarks. Returns false if bookmark with that user email
// already exists. True otherwise.

func CreateBookmark(bookmark Bookmark) bool {
	var bookmarkWithSameUserLink Bookmark
	err := config.BookmarkCollection.FindOne(context.TODO(),
		bson.D{{"email", bookmark.Email},{Key: "link" , Value: bookmark.Link}}).Decode(&bookmarkWithSameUserLink)
	if err == nil {
		// A bookmark with same user email and link already exists.
		return false
	} else {
		// A bookmark with same user email and link does not exist.
		if _, err := config.BookmarkCollection.InsertOne(context.TODO(), bookmark); err != nil {
			return false
		} else {
			return true
		}
	}
}

// Update one bookmark in which the bookmark fields email and link match,
//Returns true if the operation was successful, false otherwise
func UpdateBookmark(bookmark Bookmark) bool {

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"email", bookmark.Email},{Key: "link" , Value: bookmark.Link}}
	update := bson.M{"$set":bookmark}

	result, err := config.BookmarkCollection.UpdateOne(context.TODO(), filter, update, opts)

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return true
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		return true
	}
	log.Fatal(err)
	return false
}


//Update view count on visit
func UpdateBookmarkVisit(bookmark Bookmark) bool {

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"email", bookmark.Email},{Key: "link" , Value: bookmark.Link}}
	update := bson.D{{"$inc", bson.D{{"viewcount", 1}}}}

	result, err := config.BookmarkCollection.UpdateOne(context.TODO(), filter, update, opts)

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return true
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		return true
	}
	log.Fatal(err)
	return false
}

// delete at most one bookmark in which the bookmark fields email and link match,
//Returns true if the operation was successful, false otherwise
func DeleteBookmark(bookmark Bookmark) bool {

	// specify the SetCollation option to provide a collation that will ignore case for string comparisons
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	filter := bson.D{{"email", bookmark.Email},{Key: "link" , Value: bookmark.Link}}
	res, err := config.BookmarkCollection.DeleteOne(context.TODO(), filter , opts)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return true
}

