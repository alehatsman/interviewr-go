package vacancydb

import (
	"time"

	"github.com/atsman/interviewr-go/db/subdb"
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
	creationDate := time.Now()
	vacancy.ID = bson.NewObjectId()
	vacancy.Owner = bson.ObjectIdHex(userId)
	vacancy.CreationDate = &creationDate
	log.Debugf("db.vacancy - Create", vacancy)
	return GetVacancyC(db).Insert(vacancy)
}

func Update(db *mgo.Database, userId string, companyId string, updateModel *models.VacancyUpdateModel) (error, *models.Vacancy) {
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

func DeleteById(db *mgo.Database, userId string, id string) (error, *models.Vacancy) {
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

	err = DeleteByQuery(db, query)

	return err, &vacancy
}

func DeleteByQuery(db *mgo.Database, query map[string]interface{}) error {
	log.Debug("db.vacancy - DeleteByQuery", query)
	err, vacIds := GetIdList(db, query)
	if err != nil {
		return err
	}

	log.Debug("db.vacancy - DeleteByQuery, vacIds", vacIds)

	_, err = GetVacancyC(db).RemoveAll(query)
	if err != nil {
		return err
	}

	_, err = subdb.GetSubC(db).RemoveAll(bson.M{
		"_id": bson.M{"$in": vacIds},
	})

	return err
}

func GetIdList(db *mgo.Database, query map[string]interface{}) (error, *[]bson.ObjectId) {
	log.Debug("db.vacancy - GetIdList, query=", query)
	ids := []bson.ObjectId{}
	err := GetVacancyC(db).Find(query).Distinct("_id", &ids)
	return err, &ids
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
