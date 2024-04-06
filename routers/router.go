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

	// r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1/datasource")
	// apiv1.Use(jwt.Jwt())
	{
		apiv1.POST("/connectDB", api.ConnectDB)
		apiv1.POST("/getTableInfo", api.GetTableSchema)
	}

	return r
}
