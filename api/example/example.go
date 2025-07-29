package example

import (
	"github.com/gin-gonic/gin"
	"kubeproj.com/response"
)

type ExampleApi struct {
}

func (*ExampleApi) ExampleTest(c *gin.Context) {
	response.SuccessWithDetail(c, "Requested succeeded!", map[string]string{
		"message": "pong",
	})
}
