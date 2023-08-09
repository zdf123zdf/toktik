package main

import (
	"fmt"
	"log"
	"toktik/conf"
	"toktik/dao/db"
	"toktik/model"
	"toktik/routes"
)

func main() {
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
	//注册路由
	r := routes.InitRouter()
	// 启动8000端口
	errRun := r.Run(":8000")
	if errRun != nil {
		return
	}
}
