package main

import (
	"fmt"
	"net/http"
	"unbound/pkg/setting"
	"unbound/routers"
)

func main() {
	router := routers.InitRouter()
	//开启服务器
	//配置服务器的数据
	s := &http.Server{
		//服务器监听的网络地址
		Addr: fmt.Sprintf(":%d", 9000),
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
}
