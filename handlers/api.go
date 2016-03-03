package handlers

import (
	"github.com/atsman/interviewr-go/handlers/users"
	"github.com/atsman/interviewr-go/middlewares"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func login(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	db.C("users").Find(bson.M{"name": "test"})
}

func BuildRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Connect())
	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.Cors())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r1 := r.Group("/api/v1")
	{
		//r1.POST("/login", login)
		r1.POST("/users", users.Create)
		//r1.GET("/images/:id", getImage)
	}

	authR := r1.Group("/")
	{
		//authR.POST("/images", uploadImage)

		authR.GET("/users", users.List)
		authR.PUT("/users/:id", users.Update)
		authR.DELETE("/users/:id", users.Delete)
		/* authR.GET("/users/:id/companies", getUserCompanies)

		authR.GET("/companies", getCompaniesList)
		authR.POST("/companies", createCompany)
		authR.GET("/companies/:id", getCompany)
		authR.PUT("/companies/:id", updateCompany)
		authR.DELETE("/companies/:id", deleteCompany)

		authR.GET("/companies/:id/comments", getCompanyComments)
		authR.POST("/companies/:id/comments", createCompanyComment)

		authR.GET("/vacancies", getVacanciesList)
		authR.POST("/vacancies", createVacancy)
		authR.GET("/vacancies/:id", getVacancy)
		authR.PUT("/vacancies/:id", updateVacancy)
		authR.DELETE("/vacancies/:id", deleteVacancy)
		authR.GET("/vacancies/:id/subscription", getVacancySubscriptions)*/
	}
	return r
}
