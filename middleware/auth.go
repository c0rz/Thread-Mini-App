package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddlewareUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "There is no login access yet.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "There is no login access yet.",
			})
			c.Abort()
			return
		}
		_, lvl, err := ExtractTokenID(c)
		if err != nil {
			panic(err)
		}

		if lvl != "Admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Hello Member.",
			})
			c.Abort()
			return
		}
		
		c.Next()
	}
}
