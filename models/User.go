package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionUsers = "users"
)

type User struct {
	ID          bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string          `json:"email" bson:"email"`
	Username    string          `json:"username" bson:"username"`
	Password    string          `json:"password" bson:"password"`
	Name        string          `json:"name" bson:"name"`
	Surname     string          `json:"surname" bson:"surname"`
	About       string          `json:"about" bson:"about"`
	Сountry     string          `json:"country" bson:"country"`
	Phone       string          `json:"phone" bson:"phone"`
	Dob         time.Time       `json:"dob" bson:"dob"`
	ImageID     bson.ObjectId   `json:"imageId,omitempty" bson:"imageId,omitempty"`
	Companies   []bson.ObjectId `json:"companies,omitempty" bson:"companies,omitempty"`
	Experiences []Expirience    `json:"experiences" bson:"experiences"`
	Social      Social          `json:"social" bson:"social"`
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
