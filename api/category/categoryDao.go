package category;

import "gopkg.in/mgo.v2/bson"
import "github.com/superbarne/drueckMich/api/database"

const (
	COLLECTION = "category"
)

func Find(query bson.M, sort string) ([]Category, error) {
	var users []Category
	err := database.DB.C(COLLECTION).Find(query).Sort(sort).All(&users)
	return users, err
}

func Get(id string) (Category, error) {
	var category Category
	err := database.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&category)
	return category, err
}

func Create(category Category) error {
	err := database.DB.C(COLLECTION).Insert(&category)
	return err
}

func RemoveMany(query bson.M) error {
	_, err := database.DB.C(COLLECTION).RemoveAll(query)
	return err
}

func Remove(category Category) error {
	err := database.DB.C(COLLECTION).Remove(&category)
	return err
}

func Update(category Category) error {
	err := database.DB.C(COLLECTION).UpdateId(category.ID, &category)
	return err
}
