// Package middlewares gin 中间件
package middlewares

import (
	"errors"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// ForceUA 强制请求附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取 User-Agent 信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"),
				"请求必须附带 User-Agent 标头")
			return
		}
		c.Next()
	}
}
