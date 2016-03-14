package subscriptions

import (
	"net/http"

	"github.com/atsman/interviewr-go/db/subdb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("handlers.subscriptions")

func getSub(c *gin.Context) (error, *models.Subscription) {
	sub := models.Subscription{}
	err := c.Bind(&sub)
	return err, &sub
}

func Create(c *gin.Context) {
	db := utils.GetDb(c)
	userId := utils.GetUserId(c)

	err, sub := getSub(c)
	if err != nil {
		log.Error("Binding error", err)
		c.Error(err)
		return
	}

	err = subdb.Create(db, userId, sub)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, sub)
}

func Delete(c *gin.Context) {
	db := utils.GetDb(c)
	userId := utils.GetUserId(c)
	id := c.Params.ByName("id")
	err, sub := subdb.Delete(db, userId, id)
	if err != nil {
		log.Debug(err)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sub)
}

func GetOne(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")
	err, sub := subdb.GetOne(db, id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sub)
}

func GetList(c *gin.Context) {
	db := utils.GetDb(c)
	query := BuildQuery(c)
	err, subs := subdb.GetList(db, query)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, subs)
}
