package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionInterviews = "interviews"
)

const (
	NotStarted = "NotStarted"
	InProgress = "InProgress"
	Completed  = "Completed"
)

type Message struct {
	UserId    string `json:"userId"`
	Username  string `json:"username"`
	Time      string `json:"time"`
	Message   string `json:"message"`
	RoomID    string `json:"roomId"`
	UserImage string `json:"userImage"`
}

type Interview struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Date      time.Time     `json:"date" bson:"date" binding:"required"`
	Owner     bson.ObjectId `json:"owner" bson:"owner"`
	Vacancy   bson.ObjectId `json:"vacancy" bson:"vacancy" binding:"required"`
	Candidate bson.ObjectId `json:"candidate" bson:"candidate" binding:"required"`
	Company   bson.ObjectId `json:"company" bson:"company" binding:"required"`
	Feedback  *Feedback     `json:"feedback,omitempty" bson:"feedback,omitempty"`
	Status    string        `json:"status,omitempty" bson:"status"`
	Duration  int           `json:"duration,omitempty" bson:"duration,omitempty"`
	Messages  *[]Message    `json:"messages,omitempty" bson:"messages,omitempty"`
}

type InterviewUpdateModel struct {
	Title     string        `json:"title,omitempty" bson:"title,omitempty"`
	Date      time.Time     `json:"date,omitempty" bson:"date,omitempty"`
	Vacancy   bson.ObjectId `json:"vacancy,omitempty" bson:"vacancy,omitempty"`
	Candidate bson.ObjectId `json:"candidate,omitempty" bson:"candidate,omitempty"`
	Company   bson.ObjectId `json:"company,omitempty" bson:"company,omitempty"`
	Status    string        `json:"status,omitempty" bson:"status,omitempty"`
	Duration  int           `json:"duration,omitempty" bson:"duration,omitempty"`
}

type InterviewViewModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title,omitempty" bson:"title"`
	Date      *time.Time    `json:"date,omitempty" bson:"date"`
	Owner     *User         `json:"owner,omitempty" bson:"owner"`
	Vacancy   *Vacancy      `json:"vacancy,omitempty" bson:"vacancy" binding:"required"`
	Candidate *User         `json:"candidate,omitempty" bson:"candidate" binding:"required"`
	Company   *Company      `json:"company,omitempty" bson:"company" binding:"required"`
	Feedback  Feedback      `json:"feedback,omitempty" bson:"feedback,omitempty"`
	Status    string        `json:"status,omitempty" bson:"status"`
	Duration  int           `json:"duration,omitempty" bson:"duration,omitempty"`
}

type Feedback struct {
	Comment   string            `json:"comment" bson:"comment"`
	Personal  *[]InterviewSkill `json:"personal,omitempty" bson:"personal,omitempty"`
	Technical *[]InterviewSkill `json:"technical,omitempty" bson:"technical,omitempty"`
}

type InterviewSkill struct {
	Skill string `json:"skill" bson:"skill"`
	Score int    `json:"score" bson:"score"`
}
