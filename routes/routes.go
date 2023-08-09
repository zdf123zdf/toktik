package routes

import (
	"github.com/gin-gonic/gin"
	"toktik/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 主路由组
	douyinGroup := r.Group("/douyin")
	{
		// feed视频流
		douyinGroup.GET("/feed/", controller.GetFeed)
	}
	return r
}
