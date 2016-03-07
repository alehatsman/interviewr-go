package vacancydb

import (
	"time"

	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.vacancy")

func GetVacancyC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionVacancies)
}

func Create(db *mgo.Database, userId string, vacancy *models.Vacancy) error {
	vacancy.ID = bson.NewObjectId()
	vacancy.Owner = bson.ObjectIdHex(userId)
	vacancy.CreationDate = time.Now()
	return GetVacancyC(db).Insert(vacancy)
}

func Update(db *mgo.Database, userId string, companyId string, updateModel *map[string]interface{}) (error, *models.Vacancy) {
	updatedVacancy := models.Vacancy{}
	hUserID := bson.ObjectIdHex(userId)
	hVacancyID := bson.ObjectIdHex(companyId)

	err := GetVacancyC(db).Update(bson.M{
		"_id":   hVacancyID,
		"owner": hUserID,
	}, bson.M{
		"$set": updateModel,
	})

	if err != nil {
		return err, &updatedVacancy
	}

	err = GetVacancyC(db).FindId(hVacancyID).One(&updatedVacancy)
	return err, &updatedVacancy
}

func Delete(db *mgo.Database, userId string, id string) (error, *models.Vacancy) {
	hUserID := bson.ObjectIdHex(userId)
	hVacancyId := bson.ObjectIdHex(id)

	query := bson.M{
		"_id":   hVacancyId,
		"owner": hUserID,
	}

	vacancy := models.Vacancy{}

	err := GetVacancyC(db).Find(query).One(&vacancy)
	if err != nil {
		return err, &vacancy
	}

	err = GetVacancyC(db).Remove(query)
	if err != nil {
		return err, &vacancy
	}

	return nil, &vacancy
}

func GetList(db *mgo.Database, query interface{}) (error, *[]models.Vacancy) {
	vacancies := []models.Vacancy{}
	err := GetVacancyC(db).Find(query).All(&vacancies)
	return err, &vacancies
}

func GetOne(db *mgo.Database, id string) (error, *models.Vacancy) {
	vacancy := models.Vacancy{}
	hId := bson.ObjectIdHex(id)
	err := GetVacancyC(db).FindId(hId).One(&vacancy)
	return err, &vacancy
}
