package interviews

import (
	"errors"

	"github.com/atsman/interviewr-go/commons/strutils"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var idsProps = []string{"owner", "vacancy", "candidate", "company"}

func buildRelatedToQuery(relatedTo string) []bson.M {
	hRelatedTo := bson.ObjectIdHex(relatedTo)
	return []bson.M{
		bson.M{"owner": hRelatedTo},
		bson.M{"candidate": hRelatedTo},
	}
}

func BuildQuery(c *gin.Context) (error, interface{}) {
	values := c.Request.URL.Query()
	query := map[string]interface{}{}

	for _, prop := range idsProps {
		err := utils.ProcessAndAddIfExist(prop, &values, query, utils.ConvertToObjectId)
		if err != nil {
			return err, query
		}
	}

	relatedTo := values.Get("relatedTo")

	if strutils.IsNotEmpty(relatedTo) {
		if !bson.IsObjectIdHex(relatedTo) {
			return errors.New("relatedTo is not valid ObjectIdHex"), query
		}
		query["$or"] = buildRelatedToQuery(relatedTo)
	}

	return nil, query
}
