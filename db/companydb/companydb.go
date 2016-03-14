package companydb

import (
	"github.com/atsman/interviewr-go/db/vacancydb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("db.company")

var pipeline = []bson.M{
	{"$lookup": bson.M{
		"from":         "users",
		"localField":   "owner",
		"foreignField": "_id",
		"as":           "ownerObj",
	}},

	{"$project": bson.M{
		"_id":               1,
		"name":              1,
		"category":          1,
		"description":       1,
		"owner":             utils.FirstElem("$ownerObj"),
		"short_description": 1,
		"imageId":           1,
		"location":          1,
		"email":             1,
		"phone":             1,
		"site":              1,
		"creationDate":      1,
		"specializations":   1,
	}},
}

func GetCompanyC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionCompanies)
}

func Create(db *mgo.Database, userId string, company *models.Company) error {
	company.ID = bson.NewObjectId()
	company.Owner = bson.ObjectIdHex(userId)
	return GetCompanyC(db).Insert(company)
}

func Update(db *mgo.Database, userId string, companyId string, updateModel *models.CompanyUpdateModel) (error, *models.Company) {
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

func Delete(db *mgo.Database, userId string, companyId string) (error, *models.Company) {
	hUserID := bson.ObjectIdHex(userId)
	hCompanyID := bson.ObjectIdHex(companyId)

	query := bson.M{
		"_id":   hCompanyID,
		"owner": hUserID,
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

	err = vacancydb.DeleteByQuery(db, bson.M{
		"company_id": hCompanyID,
	})

	if err != nil {
		return err, &company
	}

	return nil, &company
}

func GetOne(db *mgo.Database, id string) (error, *models.CompanyViewModel) {
	hID := bson.ObjectIdHex(id)
	company := models.CompanyViewModel{}

	findByIdAndJoin := append([]bson.M{
		{"$match": bson.M{"_id": hID}},
	}, pipeline...)

	err := GetCompanyC(db).Pipe(findByIdAndJoin).One(&company)
	return err, &company
}

func GetList(db *mgo.Database, query interface{}) (error, *[]models.CompanyViewModel) {
	companies := []models.CompanyViewModel{}

	findByAndJoin := append([]bson.M{
		{"$match": query},
	}, pipeline...)

	err := GetCompanyC(db).Pipe(findByAndJoin).All(&companies)
	return err, &companies
}
