package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionVacancies = "vacancies"
)

type Vacancy struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Company        bson.ObjectId
	Owner          bson.ObjectId
	Title          string
	CreationDate   time.Time
	Location       string
	Type           string
	Position       string
	Description    string
	RequiredSkills []string
	Subscriptions  []bson.ObjectId
}
