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

func bindComment(c *gin.Context) (error, *models.Comment) {
	comment := models.Comment{}
	err := c.Bind(&comment)
	return err, &comment
}

func Create(c *gin.Context) {
	log.Debug("companies.Create")
	err, company := bindCompany(c)
	log.Debug("companies.Create - bindCompany", err, company)
	if err != nil {
		c.JSON(http.StatusBadRequest, notValidModel(err))
		return
	}

	db := utils.GetDb(c)
	userId := utils.GetUserId(c)
	err = companydb.Create(db, userId, company)
	log.Debugf("companies.Create - userId := %v", userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, notValidModel(err))
		return
	}

	c.JSON(http.StatusCreated, company)
}

func Update(c *gin.Context) {
	updateModel := models.CompanyUpdateModel{}
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

func AddComment(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")
	err, comment := bindComment(c)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = companydb.AddComment(db, id, comment)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func GetComments(c *gin.Context) {
	db := utils.GetDb(c)
	id := c.Params.ByName("id")

	err, comments := companydb.GetComments(db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, comments)
}

func DeleteComment(c *gin.Context) {
	db := utils.GetDb(c)
	companyId := c.Params.ByName("companyId")
	commentId := c.Params.ByName("commentId")

	err := companydb.DeleteComment(db, companyId, commentId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
