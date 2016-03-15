package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionVacancies = "vacancies"
)

type Skill struct {
	Text string `json:"text" bson:"text"`
}

type Vacancy struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Company        bson.ObjectId `json:"company_id,omitempty" bson:"company_id"`
	Owner          bson.ObjectId `json:"owner,omitempty" bson:"owner"`
	Title          string        `json:"title,omitempty" bson:"title"`
	CreationDate   *time.Time    `json:"creatingDate,omitempty" bson:"creatingDate"`
	Location       string        `json:"location,omitempty" bson:"location"`
	Type           string        `json:"type,omitempty" bson:"type"`
	Position       string        `json:"position,omitempty" bson:"position"`
	Description    string        `json:"description,omitempty" bson:"description"`
	RequiredSkills *[]Skill      `json:"required_skills,omitempty" bson:"requiredSkills"`
}

type VacancyUpdateModel struct {
	Company        bson.ObjectId `json:"company_id,omitempty" bson:"company_id"`
	Title          string        `json:"title,omitempty" bson:"title"`
	CreationDate   *time.Time    `json:"creatingDate,omitempty" bson:"creatingDate"`
	Location       string        `json:"location,omitempty" bson:"location"`
	Type           string        `json:"type,omitempty" bson:"type"`
	Position       string        `json:"position,omitempty" bson:"position"`
	Description    string        `json:"description,omitempty" bson:"description"`
	RequiredSkills *[]Skill      `json:"required_skills,omitempty" bson:"requiredSkills"`
}
