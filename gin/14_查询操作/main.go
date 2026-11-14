package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, _ := gorm.Open("mysql", "root:860625@tcp(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	// 数据库迁移
	db.AutoMigrate(&User{})

	// 创建
	//u1 := User{Name: "aili", Age: 20}
	//db.Create(&u1)
	//u2 := User{Name: "zhangjunhao", Age: 6}
	//db.Create(&u2)

	// 1.查询
	var user User
	db.First(&user)
	//fmt.Println(user)
	// 2. 查询所有
	var users []User
	db.Debug().Find(&users) // Debug可以查询sql语句
	//fmt.Println(users)
	// 3. 条件查询
	var user2 User
	db.Where(&User{Name: "zl", Age: 40}).First(&user2)
	//fmt.Println(user2)

	// 4. where查询多条语句
	var user3 []User
	db.Debug().Where("name=?", "zl").Find(&user3)
	fmt.Println(user3)

	// 5. select 查询多条语句
	var user4 []User
	//db.Debug().Select("*").Find(&user4)
	db.Debug().Select("name", "age").Find(&user4)
	fmt.Println(user4)
}
