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
		// 注册接口
		douyinGroup.POST("/user/register/", controller.UserRegister)
		// 登录接口
		douyinGroup.POST("/user/login/", controller.UserLogin)
		// 用户接口
		douyinGroup.GET("/user/", controller.UserInfo)
		// 视频投稿
		douyinGroup.POST("/publish/action/", controller.PublishAction)
		// 视频发布列表
		douyinGroup.GET("/publish/list/", controller.PublishList)
	}
	return r
}
