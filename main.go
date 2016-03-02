package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}

func CreateJwtToken() {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["foo"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(mySigningKey)
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func main() {
	router := gin.Default()

	router.Use(Auth("secret"))

	router.GET("test", testHandler)

	router.Run(":8080")
}
