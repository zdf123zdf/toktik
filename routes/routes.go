package routes

import (
	"github.com/gin-gonic/gin"
	"toktik/controller"
	"toktik/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 主路由组
	toktikGroup := r.Group("/douyin")
	{
		// feed视频流
		toktikGroup.GET("/feed/", controller.GetFeed)
		// user路由组
		userGroup := toktikGroup.Group("/user")
		{
			// 用户接口
			userGroup.GET("/", middleware.AuthMiddleware(), controller.UserInfo)
			// 注册接口
			userGroup.POST("/register/", controller.UserRegister)
			// 登录接口
			userGroup.POST("/login/", controller.UserLogin)
		}
		// publish路由组
		publishGroup := toktikGroup.Group("/publish")
		{
			// 视频投稿
			publishGroup.POST("/action/", middleware.AuthMiddleware(), controller.PublishAction)
			// 视频发布列表
			publishGroup.GET("/list/", middleware.AuthMiddleware(), controller.PublishList)
		}
		// favorite路由组
		favoriteGroup := toktikGroup.Group("/favorite")
		{
			// 赞操作
			favoriteGroup.POST("/action", middleware.AuthMiddleware(), controller.FavoriteAction)
			// 喜欢列表
			favoriteGroup.GET("/list", middleware.AuthMiddleware(), controller.FavoriteList)

		}
		// comment 路由组
		commentGroup := toktikGroup.Group("/comment")
		{
			// 评论操作
			commentGroup.POST("/action", middleware.AuthMiddleware(), controller.CommentAction)
			// 视频评论列表
			commentGroup.GET("/list", middleware.AuthMiddleware(), controller.CommentList)
		}

	}
	return r
}
