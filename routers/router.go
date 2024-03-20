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
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
