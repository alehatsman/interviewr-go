package users

import (
	"errors"
	"net/http"

	"github.com/atsman/interviewr-go/db/userdb"
	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("handlers.users")

func getDb(c *gin.Context) *mgo.Database {
	db := c.MustGet("db").(*mgo.Database)
	return db
}

func getUserC(db *mgo.Database) *mgo.Collection {
	return db.C(models.CollectionUsers)
}

func getUser(c *gin.Context) (error, models.User) {
	user := models.User{}
	err := c.Bind(&user)
	return err, user
}

func Create(c *gin.Context) {
	db := getDb(c)

	err, user := getUser(c)
	if err != nil {
		c.Error(err)
		return
	}

	err = userdb.Create(db, user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	db := getDb(c)
	id := c.Params.ByName("id")

	log.Debugf("Update, Id=%v", id)

	if !bson.IsObjectIdHex(id) {
		log.Debug("Update, id is not a ObjectIdHex")
		c.Error(errors.New("id is not a ObjectId"))
		return
	}

	var user map[string]interface{}
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(err)
		return
	}

	hexID := bson.ObjectIdHex(id)

	err, updatedUser := userdb.Update(db, hexID, user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func Delete(c *gin.Context) {
	db := getDb(c)
	id := c.Params.ByName("id")
	hId := bson.ObjectIdHex(id)

	err, user := userdb.Delete(db, hId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func List(c *gin.Context) {
	db := getDb(c)

	err, users := userdb.List(db, bson.M{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}
