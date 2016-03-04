package companies

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

var log = logging.MustGetLogger("handlers.companies")

func getCompany(c *gin.Context) (error, &models.Company)  {
  company := models.Company{}
  err := c.Bind(&company)
  return err, &company
}

func Create(c *gin.Context)  {
  db := utils.GetDb(c)

  err, company := getCompany(c)
  if err != nil {
		c.Error(err)
		return
	}



}

