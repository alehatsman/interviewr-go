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

func Create(db *mgo.Database, user models.User) error {
	err := getUserC(db).Insert(user)
	return err
}

func Update(db *mgo.Database, id bson.ObjectId, user map[string]interface{}) (error, models.User) {
	err := getUserC(db).UpdateId(id, bson.M{
		"$set": user,
	})
	if err != nil {
		return err, nil
	}
	var updatedUser = models.User{}
	err = getUserC(db).FindId(id).One(&updatedUser)
	if err != nil {
		return err, nil
	}
	return err, updatedUser
}
