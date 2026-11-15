package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dsn := "root:860625@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql", dsn)

	// 连接成功后注册关闭操作
	defer db.Close()
	// 表关联
	db.AutoMigrate(&User{})

	var user User

	// 软删除
	db.Where("Name=?", "大王子2").First(&user)
	db.Debug().Delete(&user)

	// 物理删除
	db.Unscoped().Where("Name=?", "小王子").First(&user)
	db.Unscoped().Delete(&user)
}
