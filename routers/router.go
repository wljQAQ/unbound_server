package routers

import (
	"unbound/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//创建一个不带有任何中间件的

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.POST("/datasource", api.ConnectDB)

	// r.GET("/auth", api.GetAuth)

	// apiv1 := r.Group("/api/v1")
	// // apiv1.Use(jwt.Jwt())
	// {
	// 	//获取标签列表
	// 	apiv1.GET("/tags", v1.GetTags)
	// 	//新建标签
	// 	apiv1.POST("/tags", v1.AddTag)
	// 	//测试
	// 	apiv1.POST("/model", v1.ConnectDB)
	// 	//更新指定标签
	// 	apiv1.PUT("/tags/:id", v1.EditTag)
	// 	//删除指定标签
	// 	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	// 	//获取文章列表
	// 	apiv1.GET("/articles", v1.GetArticles)
	// 	//获取指定文章
	// 	apiv1.GET("/articles/:id", v1.GetArticle)
	// 	//新建文章
	// 	apiv1.POST("/articles", v1.AddArticle)
	// 	//更新指定文章
	// 	apiv1.PUT("/articles/:id", v1.EditArticle)
	// 	//删除指定文章
	// 	apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	// }

	return r
}
