// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(route *gin.Engine) {

	//注册全局中间件
	registerGlobalMiddleWare(route)

	//注册 API 路由
	routes.RegisterAPIRoutes(route)

	//配置 404 路由
	setup404Handler(route)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func setup404Handler(r *gin.Engine) {

	//处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			//默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    "404",
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
