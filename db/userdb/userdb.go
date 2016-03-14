package userdb

import (
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.user")

func GetUserC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionUsers)
}

func Create(db *mgo.Database, user *models.User) error {
	user.ID = bson.NewObjectId()
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(pass[:])
	return GetUserC(db).Insert(user)
}

func Update(db *mgo.Database, id string, user *models.User) (error, *models.User) {
	updatedUser := models.User{}
	hId := bson.ObjectIdHex(id)
	err := GetUserC(db).UpdateId(hId, bson.M{
		"$set": user,
	})
	if err != nil {
		return err, &updatedUser
	}

	err = GetUserC(db).FindId(hId).One(&updatedUser)
	return err, &updatedUser
}

func Delete(db *mgo.Database, id string) (error, *models.User) {
	user := models.User{}
	hID := bson.ObjectIdHex(id)
	err := GetUserC(db).FindId(hID).One(&user)
	if err != nil {
		return err, &user
	}

	err = GetUserC(db).RemoveId(hID)
	if err != nil {
		return err, &user
	}

	return nil, &user
}

func GetOne(db *mgo.Database, query interface{}) (error, *models.User) {
	user := models.User{}
	err := GetUserC(db).Find(query).One(&user)
	return err, &user
}

func GetOneById(db *mgo.Database, id string) (error, *models.User) {
	hID := bson.ObjectIdHex(id)
	return GetOne(db, bson.M{
		"_id": hID,
	})
}

func GetList(db *mgo.Database, query *bson.M) (error, *[]models.User) {
	users := []models.User{}
	err := GetUserC(db).Find(bson.M{}).All(&users)
	return err, &users
}
