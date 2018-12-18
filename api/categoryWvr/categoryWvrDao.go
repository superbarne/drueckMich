package categoryWvr;

import "gopkg.in/mgo.v2/bson"
import "github.com/superbarne/drueckMich/api/database"

const (
	COLLECTION = "category-wvr"
)

func Find(query bson.M, sort string) ([]CategoryWvr, error) {
	var users []CategoryWvr
	err := database.DB.C(COLLECTION).Find(query).Sort(sort).All(&users)
	return users, err
}

func Get(id string) (CategoryWvr, error) {
	var category CategoryWvr
	err := database.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&category)
	return category, err
}

func Create(category CategoryWvr) error {
	err := database.DB.C(COLLECTION).Insert(&category)
	return err
}

func RemoveMany(query bson.M) error {
	_, err := database.DB.C(COLLECTION).RemoveAll(query)
	return err
}

func Remove(category CategoryWvr) error {
	err := database.DB.C(COLLECTION).Remove(&category)
	return err
}

func Update(category CategoryWvr) error {
	err := database.DB.C(COLLECTION).UpdateId(category.ID, &category)
	return err
}
