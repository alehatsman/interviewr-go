package vacancies

import (
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/gin-gonic/gin"
)

func BuildQuery(c *gin.Context) interface{} {
	values := c.Request.URL.Query()
	query := map[string]interface{}{}

	utils.ProcessAndAddIfExist("company_id", &values, query, utils.ConvertToObjectId)
	utils.ProcessAndAddIfExist("owner", &values, query, utils.ConvertToObjectId)

	return query
}
