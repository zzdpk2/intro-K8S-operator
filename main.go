package main

import (
	"kubeproj.com/global"
	"kubeproj.com/initializer"
)

func main() {
	r := initializer.Routers()
	initializer.Viper()
	initializer.KubernetesClient()
	panic(r.Run(global.CONF.System.Addr))
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
