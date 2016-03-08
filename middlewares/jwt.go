package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ID      = "_id"
	USER_ID = "USER_ID"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		userId := token.Claims[ID]

		c.Set(USER_ID, userId)
		c.Next()
	}
}
