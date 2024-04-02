// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goblog/routes"
	"strings"
)

func SetupRoute(router *gin.Engine) {

	//注册全局中间件
	registerGlobalMiddleware(router)

	//注册 API 路由
	routes.RegisterAPIRoutes(router)

	//配置 404 路由
	setup404Notfound(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Notfound(router *gin.Engine) {
	//处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		//获取标头 Accept 信息
		acceptStr := c.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			//如果是 HTML
			c.String(404, "页面返回 404")
		} else {
			c.JSON(404, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
