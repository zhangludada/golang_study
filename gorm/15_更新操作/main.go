package main

import (
	"fmt"

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

	// 更新方法一
	// 默认会修改所有字段
	db.First(&user)
	user.Age = 26
	db.Save(&user)

	// 方法二
	db.Model(&user).Update("age", 20)

	// 方法三
	db.Model(&user).Updates(map[string]interface{}{"name": "jinzhu", "age": 25})

	// 方法四
	db.Model(&user).Updates(User{Name: "小王子", Age: 25})

	fmt.Println(user)
}
