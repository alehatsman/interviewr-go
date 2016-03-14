package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionSubscriptions = "subscription"
)

type Subscription struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Vacancy   bson.ObjectId `json:"vacancy,omitempty" bson:"vacancy,omitempty"`
	Candidate bson.ObjectId `json:"candidate,omitempty" bson:"candidate,omitempty"`
	Interview bson.ObjectId `json:"interview,omitempty" bson:"interview,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type SubscriptionViewModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Vacancy   Vacancy       `json:"vacancy,omitempty" bson:"vacancy,omitempty"`
	Candidate User          `json:"candidate,omitempty" bson:"candidate,omitempty"`
	Interview Interview     `json:"interview,omitempty" bson:"interview,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
