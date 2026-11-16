package main

import (
	"main/dao"
	"main/models"
	"main/routers"
)

func main() {
	// 链接数据库
	dao.InitMysql()
	// 数据库迁移
	dao.DB.AutoMigrate(&models.Todo{})
	// 路由
	r := routers.SetUpRouter()
	r.Run(":8080")
}
