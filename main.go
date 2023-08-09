package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"log"
	"toktik/conf"
	"toktik/controller"
	"toktik/dao/db"
	"toktik/model"
)

func main() {
	router := gin.Default()
	// 初始化配置文件
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
	// 初始化数据库连接
	err = db.InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	// 自动迁移表
	err = db.DB.AutoMigrate(&model.Video{})
	if err != nil {
		log.Fatal("迁移数据库失败:", err)
	}
	ginpprof.Wrapper(router)
	router.GET("/douyin/feed/", controller.GetFeed)
	_ = router.Run(":8000")
}
