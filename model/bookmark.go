package model



import (
"MarkMyLink/config"
"go.mongodb.org/mongo-driver/bson"
//"go.mongodb.org/mongo-driver/mongo"
"context"
"log"
"fmt"
//"gopkg.in/mgo.v2/bson"
)

// BookMark struct()
type BookMarkCat struct {
	Name       string        `bson:"Name"`
	Link       string        `bson:"Link"`
	Viewcount  string			 `bson:"Viewcount"`
	Timestamp  string        `bson:"Timestamp"`
}

func AllBookMarks() ([]BookMarkCat, error) {

	//var bm []BookMarkCat
	bm := []BookMarkCat{}
	bmCursor, err := config.BookmarkCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = bmCursor.All(context.TODO(), &bm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hi")
	fmt.Println(bm)
	return bm, nil
}
