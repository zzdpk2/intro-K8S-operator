package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, "+
		"Token,X-Token,X-User-Id")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
		"Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
	c.Header("Access-Control-Allow-Credentials", "true")

	// Release all OPTIONS methods
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}

	c.Next()
}
