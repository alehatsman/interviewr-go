package vacancies

import (
	"net/http"

	"github.com/atsman/interviewr-go/db/subdb"
	"github.com/atsman/interviewr-go/db/vacancydb"
	"github.com/atsman/interviewr-go/handlers/utils"
	"github.com/atsman/interviewr-go/models"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("handlers.vacancies")

var badRequestError = utils.ApiError{
	Status: http.StatusBadRequest,
	Title:  "Vacancy model not valid",
}

var notFoundError = utils.ApiError{
	Status: http.StatusNotFound,
	Title:  "Vacancy not found",
}

func getVacancy(c *gin.Context) (error, *models.Vacancy) {
	vacancy := models.Vacancy{}
	err := c.Bind(&vacancy)
	return err, &vacancy
}

func Create(c *gin.Context) {
	db := utils.GetDb(c)
	userId := utils.GetUserId(c)

	err, vacancy := getVacancy(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, badRequestError)
		return
	}

	err = vacancydb.Create(db, userId, vacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, badRequestError)
		return
	}

	c.JSON(http.StatusCreated, vacancy)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")

	updateModel := map[string]interface{}{}
	err := c.BindJSON(&updateModel)
	if err != nil {
		c.Error(err)
		return
	}

	db := utils.GetDb(c)
	userId := utils.GetUserId(c)
	err, updatedVacancy := vacancydb.Update(db, userId, id, &updateModel)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedVacancy)
}

func Delete(c *gin.Context) {
	db := utils.GetDb(c)
	userId := utils.GetUserId(c)
	id := c.Params.ByName("id")
	err, user := vacancydb.Delete(db, userId, id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusNotFound, notFoundError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetList(c *gin.Context) {
	db := utils.GetDb(c)
	query := BuildQuery(c)

	err, vacancies := vacancydb.GetList(db, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, vacancies)
}

func GetOne(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	err, vacancy := vacancydb.GetOne(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, notFoundError)
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

func GetVacancySubscriptions(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	query := map[string]interface{}{}
	query["vacancy"] = id
	//query["vacancy"]

	err, subs := subdb.GetList(db, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, subs)
}
