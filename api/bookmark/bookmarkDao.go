package bookmark;

import "gopkg.in/mgo.v2/bson"
import "github.com/superbarne/drueckMich/api/database"

const (
	COLLECTION = "bookmark"
)

func Find(query bson.M, sort string) ([]Bookmark, error) {
	var users []Bookmark
	err := database.DB.C(COLLECTION).Find(query).Sort(sort).All(&users)
	return users, err
}

func Get(id string) (Bookmark, error) {
	var bookmark Bookmark
	err := database.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&bookmark)
	return bookmark, err
}

func Create(bookmark Bookmark) error {
	err := database.DB.C(COLLECTION).Insert(&bookmark)
	return err
}

func Remove(bookmark Bookmark) error {
	err := database.DB.C(COLLECTION).Remove(bson.M{"_id": bookmark.ID })
	return err
}

func Update(bookmark Bookmark) error {
	err := database.DB.C(COLLECTION).UpdateId(bookmark.ID, &bookmark)
	return err
}