package subdb

import (
	"github.com/atsman/interviewr-go/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var pipeline = []bson.M{
	{"$lookup": bson.M{
		"from":         "users",
		"localField":   "candidate",
		"foreignField": "_id",
		"as":           "candidate",
	}},

	{"$lookup": bson.M{
		"from":         "interviews",
		"localField":   "interview",
		"foreignField": "_id",
		"as":           "interview",
	}},

	{"$project": bson.M{
		"_id":       1,
		"candidate": 1,
		"interview": 1,
		"createdAt": 1,
	}},
}

func GetSubC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionSubscriptions)
}

func Create(db *mgo.Database, sub *models.Subscription) error {
	sub.ID = bson.NewObjectId()
	return GetSubC(db).Insert(sub)
}

func Delete(db *mgo.Database, userId string, id string) (error, *models.Subscription) {
	query := bson.M{
		"_id":       id,
		"candidate": userId,
	}

	sub := models.Subscription{}
	err := GetSubC(db).FindId(query).One(&sub)
	if err != nil {
		return err, &sub
	}

	err = GetSubC(db).Remove(query)
	return err, &sub
}

func GetOne(db *mgo.Database, userId string) (error, *models.SubscriptionViewModel) {
	sub := models.SubscriptionViewModel{}
	err := GetSubC(db).Pipe(pipeline).One(&sub)
	return err, &sub
}
