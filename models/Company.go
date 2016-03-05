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
	Name             string        `json:"name" bson:"name" binding:"required"`
	Category         string        `json:"category" bson:"category" binding:"required"`
	Description      string        `json:"description" bson:"description" binding:"required"`
	Owner            bson.ObjectId `json:"owner" bson:"owner"`
	ShortDescription string        `json:"shortDescription" bson:"shortDescription" binding:"required"`
	Location         string        `json:"location" bson:"location" binding:"required"`
	Email            string        `json:"email" bson:"email" binding:"required"`
	Phone            string        `json:"phone" bson:"phone"`
	Site             string        `json:"site" bson:"site" binding:"required"`
	CreationDate     time.Time     `json:"creationDate" bson:"creationDate"`
	Specializations  []string      `json:"specializations" bson:"specializations"`
}
