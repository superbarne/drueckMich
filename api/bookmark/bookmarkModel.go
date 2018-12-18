package bookmark;

import "gopkg.in/mgo.v2/bson"
import "time"

type Bookmark struct {
	ID          		bson.ObjectId 	`bson:"_id" json:"id"`
	UserId					bson.ObjectId 	`bson:"userId" json:"userId"`
	CategoryIds			[]bson.ObjectId `bson:"categoryIds" json:"categoryIds"`
	CategoryWvrIds	[]bson.ObjectId `bson:"categoryWvrIds" json:"categoryWvrIds"`
	Title						string        	`bson:"title" json:"title"`
	ImageUrls				[]string        `bson:"images" json:"images"`
	Url							string        	`bson:"url" json:"url"`
	IconUrl					string        	`bson:"icon" json:"icon"`
	Description			string        	`bson:"description" json:"description"`
	CreatedAt				time.Time       `bson:"createdAt" json:"createdAt"`
}