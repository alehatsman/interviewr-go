package interviewdb

import (
	"github.com/atsman/interviewr-go/db/subdb"
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.interviewdb")

func GetInterviewC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionInterviews)
}

func Create(db *mgo.Database, userId string, interview *models.Interview) error {
	interview.ID = bson.NewObjectId()
	interview.Owner = bson.ObjectIdHex(userId)
	return GetInterviewC(db).Insert(interview)
}

func Update(db *mgo.Database, userId string, interviewId string, updateModel *map[string]interface{}) (error, *models.Interview) {
	hUserID := bson.ObjectIdHex(userId)
	hInterviewID := bson.ObjectIdHex(interviewId)

	err := GetInterviewC(db).Update(bson.M{
		"_id":   hInterviewID,
		"owner": hUserID,
	}, bson.M{
		"$set": updateModel,
	})

	if err != nil {
		return err, &models.Interview{}
	}

	updatedInterview := models.Interview{}
	err = GetInterviewC(db).FindId(hInterviewID).One(&updatedInterview)
	return err, &updatedInterview
}

func Delete(db *mgo.Database, userId string, interviewId string) (error, *models.Interview) {
	hUserID := bson.ObjectIdHex(userId)
	hInterviewID := bson.ObjectIdHex(interviewId)

	interview := models.Interview{}
	findQuery := bson.M{
		"_id":   hInterviewID,
		"owner": hUserID,
	}

	err := GetInterviewC(db).Find(findQuery).One(&interview)
	if err != nil {
		return err, &interview
	}

	err = GetInterviewC(db).Remove(findQuery)
	subdb.GetSubC(db).Remove(bson.M{
		"interview": interview.ID,
	})

	return err, &interview
}
