package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400
)

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "请求参数错误",
}

// Result 函数用于封装 gin 框架的 JSON 响应
//
// 参数：
// gin：gin.Context 类型，表示 HTTP 请求的上下文对象
// code：int 类型，表示 HTTP 响应的状态码
// msg：string 类型，表示 HTTP 响应的消息
// data：interface{} 类型，表示 HTTP 响应的数据
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// Ok 函数用于向客户端返回成功响应
// 参数c为gin.Context类型，表示HTTP请求上下文
// 参数data为interface{}类型，表示要返回给客户端的数据
func Ok(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, MsgFlags[SUCCESS], data)
}

// Error 函数用于处理gin框架中的错误请求，返回错误信息给客户端
// 参数c是gin框架中的上下文对象
// 参数message是错误信息字符串
func Error(c *gin.Context, message string) {
	Result(c, ERROR, message, nil)
	fmt.Println(message)
}
