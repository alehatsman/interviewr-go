package subscriptions

import (
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/gin-gonic/gin"
)

func BuildQuery(c *gin.Context) interface{} {
	values := c.Request.URL.Query()
	query := map[string]interface{}{}

	utils.ProcessAndAddIfExist("interview", &values, query, utils.ConvertToObjectId)
	utils.ProcessAndAddIfExist("candidate", &values, query, utils.ConvertToObjectId)
	utils.ProcessAndAddIfExist("vacancy", &values, query, utils.ConvertToObjectId)

	return query
}
