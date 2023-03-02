package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weinibear/admin-server/routers/api/v1"
)


func InitRouter() *gin.Engine{
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/articles", v1.GetArticles)
		//新建标签
	}

	return r
}


