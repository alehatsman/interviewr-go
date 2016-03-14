package subdb

import (
	"time"

	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.subscriptions")

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
		"vacancy":   utils.FirstElem("$vacancyObj"),
		"candidate": utils.FirstElem("$candidateObj"),
		"interview": utils.FirstElem("$interviewObj"),
		"createdAt": 1,
	}},
}

func GetSubC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionSubscriptions)
}

func Create(db *mgo.Database, userId string, sub *models.Subscription) error {
	sub.ID = bson.NewObjectId()
	sub.CreatedAt = time.Now()
	sub.Candidate = bson.ObjectIdHex(userId)
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

func GetOne(db *mgo.Database, id string) (error, *models.SubscriptionViewModel) {
	sub := models.SubscriptionViewModel{}
	hId := bson.ObjectIdHex(id)
	findByIdAndJoin := append([]bson.M{
		{"$match": bson.M{"_id": hId}},
	}, pipeline...)
	err := GetSubC(db).Pipe(findByIdAndJoin).One(&sub)
	return err, &sub
}

func GetList(db *mgo.Database, query interface{}) (error, *[]models.SubscriptionViewModel) {
	subs := []models.SubscriptionViewModel{}
	findByAndJoin := append([]bson.M{
		{"$match": query},
	}, pipeline...)

	err := GetSubC(db).Pipe(findByAndJoin).All(&subs)
	return err, &subs
}
