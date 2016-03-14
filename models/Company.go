package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionCompanies = "companies"
)

type Specialization struct {
	Text string `json:"text" bson:"text"`
}

type Company struct {
	ID               bson.ObjectId    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string           `json:"name" bson:"name" binding:"required"`
	Category         string           `json:"category" bson:"category" binding:"required"`
	Description      string           `json:"description" bson:"description" binding:"required"`
	Owner            bson.ObjectId    `json:"owner" bson:"owner"`
	ShortDescription string           `json:"short_description" bson:"short_description" binding:"required"`
	ImageID          bson.ObjectId    `json:"imageId,omitempty" bson:"imageId,omitempty"`
	Location         string           `json:"location" bson:"location" binding:"required"`
	Email            string           `json:"email" bson:"email" binding:"required"`
	Phone            string           `json:"phone" bson:"phone"`
	Site             string           `json:"site" bson:"site" binding:"required"`
	CreationDate     time.Time        `json:"creation_date" bson:"creationDate"`
	Specializations  []Specialization `json:"specializations" bson:"specializations"`
}

type CompanyUpdateModel struct {
	Name             string           `json:"name" bson:"name" binding:"required"`
	Category         string           `json:"category" bson:"category" binding:"required"`
	Description      string           `json:"description" bson:"description" binding:"required"`
	ShortDescription string           `json:"short_description" bson:"short_description" binding:"required"`
	ImageID          bson.ObjectId    `json:"imageId,omitempty" bson:"imageId,omitempty"`
	Location         string           `json:"location" bson:"location" binding:"required"`
	Email            string           `json:"email" bson:"email" binding:"required"`
	Phone            string           `json:"phone" bson:"phone"`
	Site             string           `json:"site" bson:"site" binding:"required"`
	CreationDate     time.Time        `json:"creation_date" bson:"creationDate"`
	Specializations  []Specialization `json:"specializations" bson:"specializations"`
}

type CompanyViewModel struct {
	ID               bson.ObjectId    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string           `json:"name" bson:"name" binding:"required"`
	Category         string           `json:"category" bson:"category" binding:"required"`
	Description      string           `json:"description" bson:"description" binding:"required"`
	Owner            User             `json:"owner" bson:"owner"`
	ShortDescription string           `json:"short_description" bson:"short_description" binding:"required"`
	ImageID          bson.ObjectId    `json:"imageId,omitempty" bson:"imageId,omitempty"`
	Location         string           `json:"location" bson:"location" binding:"required"`
	Email            string           `json:"email" bson:"email" binding:"required"`
	Phone            string           `json:"phone" bson:"phone"`
	Site             string           `json:"site" bson:"site" binding:"required"`
	CreationDate     time.Time        `json:"creation_date" bson:"creationDate"`
	Specializations  []Specialization `json:"specializations" bson:"specializations"`
}
