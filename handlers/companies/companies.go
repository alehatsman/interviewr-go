package companies

import (
	"net/http"

	"github.com/atsman/interviewr-go/db/companydb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("handlers.companies")

var companyNotFoundError = utils.ApiError{
	Status: http.StatusNotFound,
	Title:  "Company with provided id not found",
}

func notValidModel(err error) *utils.ApiError {
	return &utils.ApiError{
		Status:      http.StatusBadRequest,
		Title:       "Company model not valid",
		Description: err.Error(),
	}
}

func bindCompany(c *gin.Context) (error, *models.Company) {
	company := models.Company{}
	err := c.Bind(&company)
	return err, &company
}

func Create(c *gin.Context) {
	err, company := bindCompany(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, notValidModel(err))
		return
	}

	db := utils.GetDb(c)
	userId := utils.GetUserId(c)
	err = companydb.Create(db, userId, company)
	if err != nil {
		c.JSON(http.StatusBadRequest, notValidModel(err))
		return
	}

	c.JSON(http.StatusCreated, company)
}

func Update(c *gin.Context) {
	updateModel := map[string]interface{}{}
	err := c.BindJSON(&updateModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, notValidModel(err))
		return
	}

	db := utils.GetDb(c)
	id := c.Params.ByName("id")
	userID := utils.GetUserId(c)
	err, updatedCompany := companydb.Update(db, userID, id, &updateModel)
	if err != nil {
		c.JSON(http.StatusNotFound, companyNotFoundError)
		return
	}

	c.JSON(http.StatusOK, updatedCompany)
}

func Delete(c *gin.Context) {
	db := utils.GetDb(c)
	userId := utils.GetUserId(c)
	companyId := c.Params.ByName("id")

	err, company := companydb.Delete(db, userId, companyId)
	if err != nil {
		c.JSON(http.StatusNotFound, companyNotFoundError)
		return
	}

	c.JSON(http.StatusOK, company)
}

func GetOne(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	err, company := companydb.GetOne(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, companyNotFoundError)
		return
	}

	c.JSON(http.StatusOK, company)
}

func GetList(c *gin.Context) {
	db := utils.GetDb(c)
	query := BuildQuery(c)
	err, companies := companydb.GetList(db, query)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusNotFound, companyNotFoundError)
		return
	}
	c.JSON(http.StatusOK, companies)
}
