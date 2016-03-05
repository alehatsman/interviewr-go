package utils

import (
	"github.com/atsman/interviewr-go/middlewares"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func GetDb(c *gin.Context) *mgo.Database {
	db := c.MustGet("db").(*mgo.Database)
	return db
}

func GetUserId(c *gin.Context) string {
	return c.MustGet(middlewares.USER_ID).(string)
}
