package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionInterviews = "interviews"
)

type Interview struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Date      time.Time     `json:"date" bson:"date"`
	Owner     bson.ObjectId `json:"owner" bson:"owner"`
	Vacancy   bson.ObjectId `json:"vacancy" bson:"vacancy"`
	Candidate bson.ObjectId `json:"candidate" bson:"candidate"`
	Company   bson.ObjectId `json:"company" bson:"company"`
}

type InterviewViewModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title,omitempty" bson:"title,omitempty"`
	Date      time.Time     `json:"date,omitempty" bson:"date,omitempty"`
	Owner     User          `json:"owner,omitempty" bson:"owner,omitempty"`
	Vacancy   Vacancy       `json:"vacancy,omitempty" bson:"vacancy,omitempty"`
	Candidate User          `json:"candidate,omitempty" bson:"candidate,omitempty"`
	Company   Company       `json:"company,omitempty" bson:"company,omitempty"`
}
