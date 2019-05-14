package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errResult struct {
	Status  int      `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

//操作正确的时候快速拼凑json响应请求
func okJson(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    msg,
		"data":   nil,
	})
}

func dataJson(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    msg,
		"data":   data,
	})
}

//NoResponse 请求的url不存在，返回404
func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"msg":    "404, page not exists!",
	})
}

//报错的时候快速拼凑json响应请求
func errJson(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status": 1,
		"msg": msg,
		"data": nil,
	})
}
