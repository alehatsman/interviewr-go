package interviewrgo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func errHandler(err) {
	if err != nil {
		panic(err)
	}
}

func connect() {
	session, err := mgo.Dial("localhost")
	errHandler(err)
	defer session.Close()

	c := session.DB("interviewr")
}
