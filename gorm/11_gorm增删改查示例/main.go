package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 对应数据库中的表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 链接数据库
	dsn := "root:860625@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 连接成功后注册关闭操作
	defer db.Close()

	// 自动迁移：创建表
	if err := db.AutoMigrate(&UserInfo{}).Error; err != nil {
		log.Fatal("创建表失败:", err)
	}

	// 创建数据
	u1 := UserInfo{
		Name:   "张钧皓",
		Gender: "男",
		Hobby:  "蝶泳",
	}

	// 插入数据
	if err := db.Create(&u1).Error; err != nil {
		log.Fatal("插入数据失败:", err)
	}

	log.Println("创建表和插入数据成功，ID:", u1.ID)

	// 查找数据
	var u2 UserInfo
	db.First(&u2)
	fmt.Println(u2)

	// 更新
	db.Model(&u2).Update("hobby", "frisbee")

	// 删除数据
	db.Delete(&u2)

}
