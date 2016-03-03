package middlewares

import (
	"net/http"

	"github.com/atsman/interviewr-go/db"
	"github.com/gin-gonic/gin"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := db.Session.Clone()

		defer s.Close()

		c.Set("db", s.DB(db.Mongo.Database))
		c.Next()
	}
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// TODO: Handle it in a better way
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": c.Errors,
			})
		}
	}
}
