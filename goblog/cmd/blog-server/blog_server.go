package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/bootstrap"
)

func main() {

	//初始化 gin 实例
	r := gin.New()

	//初始化路由
	bootstrap.SetupRoute(r)

	//运行服务
	err := r.Run(":8080")
	if err != nil {
		//错误打印
		fmt.Println(err.Error())
	}
}
