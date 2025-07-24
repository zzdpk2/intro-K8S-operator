package initializer

import (
	"github.com/gin-gonic/gin"
	"kubeproj.com/router"
)

func Routers() {
	r := gin.Default()
	// exampleGroup := api.ApiGroupApp.ExampleApiGroup
	// r.GET("/ping", exampleGroup.ExampleTest)
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	exampleGroup.InitExample(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
