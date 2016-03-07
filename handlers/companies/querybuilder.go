package companies

import (
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func BuildQuery(c *gin.Context) interface{} {
	values := c.Request.URL.Query()
	query := map[string]interface{}{}

	utils.ProcessAndAddIfExist("owner", &values, query, func(value interface{}) interface{} {
		strVal := value.(string)
		hexVal := bson.ObjectIdHex(strVal)
		return hexVal
	})

	log.Debugf("Values %v", query)

	return &query
}
