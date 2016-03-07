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
	Company        bson.ObjectId `json:"company" bson:"company"`
	Owner          bson.ObjectId `json:"owner" bson:"owner"`
	Title          string        `json:"title" bson:"title"`
	CreationDate   time.Time     `json:"creatingDate" bson:"creatingDate"`
	Location       string        `json:"location" bson:"location"`
	Type           string        `json:"type" bson:"type"`
	Position       string        `json:"position" bson:"position"`
	Description    string        `json:"description" bson:"description"`
	RequiredSkills []string      `json:"requiredSkills" bson:"requiredSkills"`
}
