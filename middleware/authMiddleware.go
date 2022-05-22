package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/rabbice/ecommerce/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "no authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("firstname", claims.FirstName)
		c.Set("lastname", claims.LastName)
		c.Set("uid", claims.Uid)
		c.Set("seller", claims.Seller)
		c.Next()
	}
}
