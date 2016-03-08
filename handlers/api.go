package handlers

import (
	"github.com/atsman/interviewr-go/handlers/companies"
	"github.com/atsman/interviewr-go/handlers/images"
	"github.com/atsman/interviewr-go/handlers/interviews"
	"github.com/atsman/interviewr-go/handlers/meta"
	"github.com/atsman/interviewr-go/handlers/subscriptions"
	"github.com/atsman/interviewr-go/handlers/users"
	"github.com/atsman/interviewr-go/handlers/vacancies"
	"github.com/atsman/interviewr-go/middlewares"
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Connect())
	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.Cors())

	r1 := r.Group("/api/v1")
	{
		r1.POST("/login", users.Login)
		r1.POST("/users", users.Create)
		r1.GET("/images/:id", images.GetOne)
	}

	r1.POST("/images", images.Create)

	authR := r1.Group("/")
	authR.Use(middlewares.Auth("secret"))
	{

		authR.GET("/users", users.GetList)
		authR.GET("/users/:id", users.GetOne)
		authR.PUT("/users/:id", users.Update)
		authR.DELETE("/users/:id", users.Delete)
		authR.GET("/users/:id/companies", users.GetUserCompanies)

		authR.GET("/companies", companies.GetList)
		authR.POST("/companies", companies.Create)
		authR.GET("/companies/:id", companies.GetOne)
		authR.PUT("/companies/:id", companies.Update)
		authR.DELETE("/companies/:id", companies.Delete)

		//authR.GET("/companies/:id/comments", getCompanyComments)
		//authR.POST("/companies/:id/comments", createCompanyComment)

		authR.GET("/vacancies", vacancies.GetList)
		authR.POST("/vacancies", vacancies.Create)
		authR.GET("/vacancies/:id", vacancies.GetOne)
		authR.PUT("/vacancies/:id", vacancies.Update)
		authR.DELETE("/vacancies/:id", vacancies.Delete)
		authR.GET("/vacancies/:id/subscriptions", vacancies.GetVacancySubscriptions)

		authR.POST("/interview", interviews.Create)
		authR.PUT("/interview/:id", interviews.Update)
		authR.DELETE("/interview/:id", interviews.Delete)
		authR.GET("/interview/:id", interviews.GetOne)
		authR.GET("/interview", interviews.GetList)

		authR.POST("/subscriptions", subscriptions.Create)
		authR.GET("/subscriptions/:id", subscriptions.GetOne)
		authR.GET("/subscriptions", subscriptions.GetList)
		authR.DELETE("/subscriptions/:id", subscriptions.Delete)

		authR.GET("/meta/tags", meta.GetTags)
		authR.GET("/meta/countries", meta.GetCountries)
		authR.GET("/meta/categories", meta.GetCategories)
		authR.GET("/meta/position", meta.GetPosition)
		authR.GET("/meta/vacancyType", meta.GetVacancyType)
	}
	return r
}
