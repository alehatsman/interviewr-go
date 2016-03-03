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
	Candidate bson.ObjectId
	Interview bson.ObjectId
	CreatedAt time.Time
}
