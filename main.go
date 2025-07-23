package main

import (
	"kubeproj.com/api"
)
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	exampleGroup := api.ApiGroupApp.ExampleApiGroup
	r.GET("/ping", exampleGroup.ExampleTest)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
