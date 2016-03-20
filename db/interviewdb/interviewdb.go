package interviewdb

import (
	"errors"

	"github.com/atsman/interviewr-go/db/subdb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.interviewdb")

var pipeline = []bson.M{
	{"$lookup": bson.M{
		"from":         "users",
		"localField":   "owner",
		"foreignField": "_id",
		"as":           "ownerObj",
	}},

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
		"from":         "companies",
		"localField":   "company",
		"foreignField": "_id",
		"as":           "companyObj",
	}},

	{"$project": bson.M{
		"_id":       1,
		"title":     1,
		"date":      1,
		"owner":     utils.FirstElem("$ownerObj"),
		"vacancy":   utils.FirstElem("$vacancyObj"),
		"candidate": utils.FirstElem("$candidateObj"),
		"company":   utils.FirstElem("$companyObj"),
	}},
}

func GetInterviewC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionInterviews)
}

func Create(db *mgo.Database, userId string, interview *models.Interview) error {
	interview.ID = bson.NewObjectId()
	interview.Owner = bson.ObjectIdHex(userId)

	err := GetInterviewC(db).Insert(interview)

	if err != nil {
		return err
	}

	return subdb.GetSubC(db).Update(bson.M{
		"vacancy":   interview.Vacancy,
		"candidate": interview.Candidate,
	}, bson.M{
		"$set": bson.M{
			"interview": interview.ID,
		},
	})
}

func Update(db *mgo.Database, userId string, interviewId string, updateModel *models.InterviewUpdateModel) (error, *models.InterviewViewModel) {
	hUserID := bson.ObjectIdHex(userId)
	hInterviewID := bson.ObjectIdHex(interviewId)

	err := GetInterviewC(db).Update(bson.M{
		"_id":   hInterviewID,
		"owner": hUserID,
	}, bson.M{
		"$set": updateModel,
	})

	if err != nil {
		return err, &models.InterviewViewModel{}
	}

	err, interview := GetOne(db, interviewId)
	return err, interview
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

	return err, &interview
}

func GetOne(db *mgo.Database, id string) (error, *models.InterviewViewModel) {
	interview := models.InterviewViewModel{}
	hId := bson.ObjectIdHex(id)
	findByIdAndJoin := append([]bson.M{
		{"$match": bson.M{"_id": hId}},
	}, pipeline...)
	err := GetInterviewC(db).Pipe(findByIdAndJoin).One(&interview)
	return err, &interview
}

func GetList(db *mgo.Database, query interface{}) (error, *[]models.InterviewViewModel) {
	interviews := []models.InterviewViewModel{}
	findByAndJoin := append([]bson.M{
		{"$match": query},
	}, pipeline...)

	err := GetInterviewC(db).Pipe(findByAndJoin).All(&interviews)
	return err, &interviews
}

func CreateFeedback(db *mgo.Database, interviewId string, feedback *models.Feedback) error {
	return GetInterviewC(db).UpdateId(bson.ObjectIdHex(interviewId), bson.M{
		"$set": bson.M{
			"feedback": feedback,
		},
	})
}

func GetFeedback(db *mgo.Database, interviewId string) (error, *models.Feedback) {
	feedback := models.Interview{}

	if !bson.IsObjectIdHex(interviewId) {
		return errors.New("Not found"), feedback.Feedback
	}

	err := GetInterviewC(db).FindId(bson.ObjectIdHex(interviewId)).One(&feedback)
	return err, feedback.Feedback
}
