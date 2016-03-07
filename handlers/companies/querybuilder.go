package companies

import (
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/gin-gonic/gin"
)

func BuildQuery(c *gin.Context) interface{} {
	values := c.Request.URL.Query()
	query := map[string]interface{}{}

	utils.ProcessAndAddIfExist("owner", &values, query, utils.ConvertToObjectId)

	log.Debugf("Values %v", query)

	return &query
}
