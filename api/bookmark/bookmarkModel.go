package bookmark;

import "gopkg.in/mgo.v2/bson"
import "time"

type Bookmark struct {
	ID          	bson.ObjectId 	`bson:"_id" json:"id"`
	UserId				bson.ObjectId 	`bson:"userId" json:"userId"`
	CategoryIds		[]bson.ObjectId `bson:"categoryIds" json:"categoryIds"`
	Title					string        	`bson:"title" json:"title"`
	Images				[]string        `bson:"images" json:"images"`
	Url						string        	`bson:"url" json:"url"`
	Description		string        	`bson:"description" json:"description"`
	CreatedAt			time.Time       `bson:"createdAt" json:"createdAt"`
}