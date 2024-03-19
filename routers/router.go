package routers

import (
	"unbound/pkg/setting"
	v1 "unbound/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//创建一个不带有任何中间件的
	gin.SetMode(setting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/tags", v1.GetTags)

	return r
}
