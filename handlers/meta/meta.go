package meta

import (
	"net/http"

	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getData(db *mgo.Database, name string) (error, map[string]interface{}) {
	res := map[string]interface{}{}
	err := db.C("metas").Find(bson.M{"name": name}).One(&res)
	return err, res
}

func sendMeta(name string, c *gin.Context) {
	db := utils.GetDb(c)
	err, res := getData(db, name)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetTags(c *gin.Context) {
	sendMeta("tags", c)
}

func GetCountries(c *gin.Context) {
	sendMeta("countries", c)
}

func GetCategories(c *gin.Context) {
	sendMeta("categories", c)
}

func GetPosition(c *gin.Context) {
	sendMeta("position", c)
}

func GetVacancyType(c *gin.Context) {
	sendMeta("vacancyType", c)
}
