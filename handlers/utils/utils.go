package utils

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func GetDb(c *gin.Context) *mgo.Database {
	db := c.MustGet("db").(*mgo.Database)
	return db
}
