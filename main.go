package main

import (
	"github.com/gin-gonic/gin"
	"kubeproj.com/initializer"
	"net/http"
)

func main() {
	initializer.Routers()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
