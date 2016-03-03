package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionComments = "comments"
)

type Comment struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Author    bson.ObjectId
	Text      string
	CreatedAt time.Time
}
