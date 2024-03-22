package jwt

import (
	"net/http"
	"time"
	"unbound/pkg/e"
	"unbound/pkg/util"

	"github.com/gin-gonic/gin"
)

func jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义状态码和返回数据
		var code int
		var data interface{}

		// 初始状态码设置为成功
		code = e.SUCCESS

		// 获取请求中的token
		token := c.Query("token")

		// 如果token为空
		if token == "" {
			// 设置状态码为参数错误
			code = e.INVALID_PARAMS
		} else {
			// 解析token
			claims, err := util.ParseToken(token)
			// 如果解析失败
			if err != nil {
				// 设置状态码为参数错误
				code = e.INVALID_PARAMS
				// 如果token已过期
			} else if time.Now().After(claims.ExpiresAt.Time) {
				// 设置状态码为token过期
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		// 如果状态码不为成功
		if code != e.SUCCESS {
			// 返回错误信息
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			// 终止请求处理
			c.Abort()
			return
		}

		// 继续处理下一个中间件
		c.Next()
	}
}
