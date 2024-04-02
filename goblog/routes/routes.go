// Package routes 路由注册
package routes

import "github.com/gin-gonic/gin"

// RegisterAPIRoutes 注册相关路由
func RegisterAPIRoutes(router *gin.Engine) {

	//v1 版本路由
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(200, "welcome to goBlog!")
		})
	}
}
