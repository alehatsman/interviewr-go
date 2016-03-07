package subscriptions

import (
	"net/http"

	"github.com/atsman/interviewr-go/db/subdb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
)

func getSub(c *gin.Context) (error, *models.Subscription) {
	sub := models.Subscription{}
	err := c.Bind(&sub)
	return err, &sub
}

func Create(c *gin.Context) {
	db := utils.GetDb(c)

	err, sub := getSub(c)
	if err != nil {
		c.Error(err)
		return
	}

	err = subdb.Create(db, sub)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, sub)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

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
