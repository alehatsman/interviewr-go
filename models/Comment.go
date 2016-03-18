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
	Author    CommentAuthor `json:"author" bson:"author"`
	Text      string        `json:"text" bson:"text"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

type CommentAuthor struct {
	ID      string        `json:"id,omitempty" bson:"id,omitempty"`
	Name    string        `json:"name,omitempty" bson:"name,omitempty"`
	Surname string        `json:"surname,omitempty" bson:"surname,omitempty"`
	ImageID bson.ObjectId `json:"imageId,omitempty" bson:"imageId,omitempty"`
}
