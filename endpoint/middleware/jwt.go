package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izhujiang/appengine/endpoint/common"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := common.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
