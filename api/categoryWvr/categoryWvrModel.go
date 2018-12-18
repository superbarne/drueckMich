package categoryWvr;

import "gopkg.in/mgo.v2/bson"
import "time"

type CategoryWvr struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserId			bson.ObjectId `bson:"userId" json:"userId"`
	Name				string        `bson:"name" json:"name"`
	CreatedAt		time.Time     `bson:"createdAt" json:"createdAt"`
}