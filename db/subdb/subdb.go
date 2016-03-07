package subdb

import (
	"time"

	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.subscriptions")

func firstElem(field string) *bson.M {
	return &bson.M{
		"$arrayElemAt": []interface{}{field, 0},
	}
}

var pipeline = []bson.M{
	{"$lookup": bson.M{
		"from":         "vacancies",
		"localField":   "vacancy",
		"foreignField": "_id",
		"as":           "vacancyObj",
	}},

	{"$lookup": bson.M{
		"from":         "users",
		"localField":   "candidate",
		"foreignField": "_id",
		"as":           "candidateObj",
	}},

	{"$lookup": bson.M{
		"from":         "interviews",
		"localField":   "interview",
		"foreignField": "_id",
		"as":           "interviewObj",
	}},

	{"$project": bson.M{
		"_id":       1,
		"vacancy":   firstElem("$vacancyObj"),
		"candidate": firstElem("$candidateObj"),
		"interview": firstElem("$interviewObj"),
		"createdAt": 1,
	}},
}

func GetSubC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionSubscriptions)
}

func Create(db *mgo.Database, sub *models.Subscription) error {
	sub.ID = bson.NewObjectId()
	sub.CreatedAt = time.Now()
	return GetSubC(db).Insert(sub)
}

func Delete(db *mgo.Database, userId string, id string) (error, *models.Subscription) {
	hId := bson.ObjectIdHex(id)
	hUserId := bson.ObjectIdHex(userId)

	query := bson.M{
		"_id":       hId,
		"candidate": hUserId,
	}

	sub := models.Subscription{}
	err := GetSubC(db).Find(query).One(&sub)
	if err != nil {
		log.Error(err)
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

func GetList(db *mgo.Database, query interface{}) (error, *[]models.SubscriptionViewModel) {
	subs := []models.SubscriptionViewModel{}
	findByAndJoin := append([]bson.M{
		{"$match": query},
	}, pipeline...)

	log.Debug(findByAndJoin)

	err := GetSubC(db).Pipe(findByAndJoin).All(&subs)
	return err, &subs
}
