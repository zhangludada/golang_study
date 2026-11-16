package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() {
	DB, _ := gorm.Open("mysql", "root:860625@tcp(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local")
	defer DB.Close()
}
