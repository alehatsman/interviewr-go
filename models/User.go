package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionUsers = "users"
)

type User struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string        `json:"email,omitempty" bson:"email,omitempty"`
	Username    string        `json:"username,omitempty" bson:"username,omitempty"`
	Password    string        `json:"password,omitempty" bson:"password,omitempty"`
	Name        string        `json:"name,omitempty" bson:"name,omitempty"`
	Surname     string        `json:"surname,omitempty" bson:"surname,omitempty"`
	About       string        `json:"about,omitempty" bson:"about,omitempty"`
	Сountry     string        `json:"country,omitempty" bson:"country,omitempty"`
	Phone       string        `json:"phone,omitempty" bson:"phone,omitempty"`
	Dob         time.Time     `json:"dob,omitempty" bson:"dob,omitempty"`
	ImageID     bson.ObjectId `json:"imageId,omitempty" bson:"imageId,omitempty"`
	Experiences []Expirience  `json:"experiences,omitempty" bson:"experiences,omitempty"`
	Social      Social        `json:"social,omitempty" bson:"social,omitempty"`
}

type Social struct {
	Twitter  string `json:"twitter" bson:"twitter"`
	Facebook string `json:"facebook" bson:"facebook"`
	Github   string `json:"github" bson:"github"`
	Linkedin string `json:"linkedin" bson:"linkedin"`
}

type Expirience struct {
	Organization string `json:"organization" bson:"organization"`
	Position     string `json:"position" bson:"position"`
	StartPeriod  string `json:"startPeriod" bson:"startPeriod"`
	EndPeriod    string `json:"endPeriod" bson:"endPeriod"`
	About        string `json:"about" bson:"about"`
}
