package userdb

import (
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.user")

func getUserC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionUsers)
}

func Create(db *mgo.Database, user *models.User) error {
	err := getUserC(db).Insert(user)
	return err
}

func Update(db *mgo.Database, id *bson.ObjectId, user *map[string]interface{}) (error, *models.User) {
	var updatedUser = models.User{}
	err := getUserC(db).UpdateId(id, bson.M{
		"$set": user,
	})
	if err != nil {
		return err, &updatedUser
	}

	err = getUserC(db).FindId(id).One(&updatedUser)
	if err != nil {
		return err, &updatedUser
	}
	return err, &updatedUser
}

func List(db *mgo.Database, query *bson.M) (error, *[]models.User) {
	var users []models.User
	err := getUserC(db).Find(bson.M{}).All(&users)
	return err, &users
}

func Delete(db *mgo.Database, id *bson.ObjectId) (error, *models.User) {
	var user = models.User{}
	err := getUserC(db).FindId(id).One(&user)
	if err != nil {
		return err, &user
	}

	err = getUserC(db).RemoveId(id)
	if err != nil {
		return err, &user
	}

	return nil, &user
}
