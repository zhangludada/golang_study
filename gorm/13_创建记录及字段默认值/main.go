package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型

type User struct {
	ID   int64
	Name string `gorm:"default:'zl'"` // 字段默认值
	Age  int64
}

func main() {
	db, _ := gorm.Open("mysql", "root:860625@tcp(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3. 创建记录
	u := User{Name: "", Age: 23}
	db.Create(&u)

}
