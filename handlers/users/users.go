package users

import (
	"net/http"

	"github.com/atsman/interviewr-go/db/companydb"
	"github.com/atsman/interviewr-go/db/userdb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("handlers.users")

var userNotFoundError = utils.ApiError{
	Status: http.StatusNotFound,
	Title:  "User not found",
}

func getUser(c *gin.Context) (error, *models.User) {
	user := models.User{}
	err := c.Bind(&user)
	return err, &user
}

func Create(c *gin.Context) {
	db := utils.GetDb(c)

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

	c.JSON(http.StatusCreated, user)
}

func Update(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	log.Debugf("Update, Id=%v", id)

	var user map[string]interface{}
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(err)
		return
	}

	err, updatedUser := userdb.Update(db, id, &user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func Delete(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	err, user := userdb.Delete(db, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetList(c *gin.Context) {
	db := utils.GetDb(c)

	err, users := userdb.GetList(db, &bson.M{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetOne(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	err, user := userdb.GetOne(db, id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusNotFound, userNotFoundError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserCompanies(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	err, companies := companydb.GetList(db, &bson.M{
		"owner": bson.ObjectIdHex(id),
	})

	if err != nil {
		c.JSON(http.StatusNotFound, userNotFoundError)
		return
	}

	c.JSON(http.StatusOK, companies)
}
