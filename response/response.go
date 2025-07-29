package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = iota
	fail
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  "successful",
	})
}

func SuccessWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
	})
}

func SuccessWithDetail(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
		"data": data,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  "failed",
	})
}

func FailWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
	})
}

func FailWithDetail(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
		"data": data,
	})
}
