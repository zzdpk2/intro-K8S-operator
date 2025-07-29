package initializer

import (
	"github.com/gin-gonic/gin"
	"kubeproj.com/middleware"
	"kubeproj.com/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	// exampleGroup := api.ApiGroupApp.ExampleApiGroup
	// r.GET("/ping", exampleGroup.ExampleTest)
	r.Use(middleware.Cors)
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	k8sGroup := router.RouterGroupApp.KubernetesRouterGroup
	exampleGroup.InitExample(r)
	k8sGroup.InitKubernetesSRouter(r)
	return r
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
