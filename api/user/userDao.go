package user;

import "gopkg.in/mgo.v2/bson"
import "github.com/superbarne/drueckMich/api/database"

const (
	COLLECTION = "user"
)

func Find(query bson.M) ([]User, error) {
	var users []User
	err := database.DB.C(COLLECTION).Find(query).All(&users)
	return users, err
}

func Get(id string) (User, error) {
	var user User
	err := database.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func Create(user User) error {
	err := database.DB.C(COLLECTION).Insert(&user)
	return err
}

func RemoveMany(query bson.M) error {
	_, err := database.DB.C(COLLECTION).RemoveAll(query)
	return err
}

func Remove(user User) error {
	err := database.DB.C(COLLECTION).Remove(&user)
	return err
}

func Update(user User) error {
	err := database.DB.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
