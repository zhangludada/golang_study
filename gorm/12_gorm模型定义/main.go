package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 对应数据库中的表
type UserInfo struct {
	gorm.Model
	// 自定义列名
	Name   string `gorm:"column:user_name"`
	Gender string
	Hobby  string
}

// 自定义表名
func (UserInfo) TableName() string {
	return "user_info_1"
}

func main() {
	// 连接数据库
	dsn := "root:860625@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	defer db.Close()

	// 自动迁移
	if err := db.AutoMigrate(&UserInfo{}).Error; err != nil {
		log.Fatal("自动迁移失败:", err)
	}

	log.Println("数据库迁移成功")
}
