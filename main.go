package main

import (
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
	// 自动迁移表
	err = db.DB.AutoMigrate(&model.Video{})
	if err != nil {
		log.Fatalln("迁移数据库失败:", err)
	}
	//// 初始化对象存储minio
	//err = minio.InitMinio()
	//if err != nil {
	//	log.Fatalln("minio初始化失败:", err)
	//}
	//注册路由
	r := routes.InitRouter()
	// 启动8000端口
	errRun := r.Run(":8000")
	if errRun != nil {
		return
	}
}
