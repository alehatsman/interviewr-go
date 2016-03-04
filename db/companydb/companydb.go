package companydb

import (
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.company")

func GetCompanyC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionCompanies)
}

func Create(db *mgo.Database, company *models.Company) error {
	return GetCompanyC(db).Insert(company)
}

func Update(db *mgo.Database, id *bson.ObjectId, user *map[string]interface{}) (error, *models.Company) {
	updatedCompany := models.Company{}
	err := GetCompanyC(db).UpdateId(id, bson.M{
		"$set": updatedCompany,
	})
	if err != nil {
		return err, &updatedCompany
	}
	err = GetCompanyC(db).FindId(id).One(&updatedCompany)
	return err, &updatedCompany
}

func List(db *mgo.Database, query *bson.M) (error, *[]models.Company) {
	var companies []models.Company
	err := GetCompanyC(db).Find(bson.M{}).All(&companies)
	return err, &companies
}

func Delete(db *mgo.Database, id *bson.ObjectId) (error, *models.Company) {
	company := models.Company{}
	err := GetCompanyC(db).FindId(id).One(&company)
	if err != nil {
		return err, &company
	}

	err = GetCompanyC(db).RemoveId(id)
	if err != nil {
		return err, &company
	}

	return nil, &company
}
