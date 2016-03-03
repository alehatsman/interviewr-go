package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionCompanies = "companies"
)

type Company struct {
	ID               bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Category         string
	Description      string
	Name             string
	Owner            bson.ObjectId
	ShortDescription string
	Location         string
	Email            string
	Phone            string
	Site             string
	CreationDate     time.Time
	Specialization   []string
	Vacancies        []bson.ObjectId //todo remove
	Comments         []bson.ObjectId //todo remove
}
