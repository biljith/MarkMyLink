package model



import (
//"MarkMyLink/config"
//"gopkg.in/mgo.v2/bson"
)

// BookMark struct()
type BookMarkCat struct {
	Name       string        `bson:"Name"`
	Link       string        `bson:"Link"`
	Viewcount 	int			 `bson:"Viewcount"`
	Timestamp  string        `bson:"Timestamp"`
}

func AllBookMarks() ([]BookMarkCat, error) {

	bm := []BookMarkCat{}

	//err := config.Bookmark.Find(bson.M{}).All(&bm)
	//if err != nil {
	//	return nil, err
	//}
	return bm, nil
}
