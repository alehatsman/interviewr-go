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
	Candidate bson.ObjectId `json:"candidate" bson:"candidate"`
	Interview bson.ObjectId `json:"interview" bson:"interview"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

type SubscriptionViewModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Candidate User          `json:"candidate"`
	Interview Interview     `json:"interview"`
	CreatedAt time.Time     `json:"createdAt"`
}
