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

func Create(db *mgo.Database, userId string, company *models.Company) error {
	company.ID = bson.NewObjectId()
	company.Owner = bson.ObjectIdHex(userId)
	return GetCompanyC(db).Insert(company)
}

func Update(db *mgo.Database, userId string, companyId string, updateModel *map[string]interface{}) (error, *models.Company) {
	updatedCompany := models.Company{}
	hUserID := bson.ObjectIdHex(userId)
	hCompanyID := bson.ObjectIdHex(companyId)

	err := GetCompanyC(db).Update(bson.M{
		"_id":   hCompanyID,
		"owner": hUserID,
	}, bson.M{
		"$set": updateModel,
	})
	if err != nil {
		return err, &updatedCompany
	}

	err = GetCompanyC(db).FindId(hCompanyID).One(&updatedCompany)
	return err, &updatedCompany
}

func List(db *mgo.Database, query *bson.M) (error, *[]models.Company) {
	var companies []models.Company
	err := GetCompanyC(db).Find(bson.M{}).All(&companies)
	return err, &companies
}

func Delete(db *mgo.Database, userId string, companyId string) (error, *models.Company) {
	hUserId := bson.ObjectIdHex(userId)
	hCompanyId := bson.ObjectIdHex(companyId)

	query := bson.M{
		"_id":   hCompanyId,
		"owner": hUserId,
	}

	company := models.Company{}
	err := GetCompanyC(db).Find(query).One(&company)
	if err != nil {
		return err, &company
	}

	err = GetCompanyC(db).Remove(query)
	if err != nil {
		return err, &company
	}

	return nil, &company
}
