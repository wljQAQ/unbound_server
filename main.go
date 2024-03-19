package main

import (
	"fmt"
	"net/http"
	"unbound/pkg/setting"
	"unbound/routers"
)

func main() {
	// //创建gin路由 返回 Gin 的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
	// router := gin.Default()
	// //使用路由创建一个get请求
	// router.GET("/test", func(c *gin.Context) {
	// 	//gin.H 就是一个map[string]interface{}
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })
	// fmt.Println(setting.HTTPPort)
	router := routers.InitRouter()
	//开启服务器
	//配置服务器的数据
	s := &http.Server{
		//服务器监听的网络地址
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		//处理所有的请求。如果为 nil，将使用 http.DefaultServeMux
		Handler: router,
		//请求的读操作在超时前订单最大持续事件。超时过后将会关闭对应的连接。
		ReadTimeout: setting.ReadTimeout,
		//响应的写操作在超时前的最大持续时间。超时后，将会关闭对应的连接
		WriteTimeout: setting.WriteTimeout,
		//请求的头部的最大字节数
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	// router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
