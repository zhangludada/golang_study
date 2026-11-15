package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//		请用postman测试
//	 所有功能已完成
func main() {
	// 链接数据库
	db, _ := gorm.Open("mysql", "root:860625@tcp(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	// 数据库迁移
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// 创建一个路由组
	v1Group := r.Group("/v1")
	{
		// 添加一个代办事项
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			todo.Status = true
			c.BindJSON(&todo)
			db.Create(&todo)
			c.JSON(http.StatusOK, todo)
		})

		// 查看所有代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var title []Todo
			db.Where("status=?", true).Find(&title)
			c.JSON(http.StatusOK, title)
		})

		// 查看一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			var td Todo
			id := c.Param("id")
			db.Where("id=?", id).First(&td)
			c.JSON(http.StatusOK, td)
		})

		// 删除一个代办事项
		// 仅修改 status = false
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			var td Todo
			id := c.Param("id")
			db.Where("id=?", id).First(&td)
			db.Model(&td).Update("status", false)
			c.JSON(http.StatusOK, td)
		})

		// 编辑一个代办事项

		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			var td Todo
			id := c.Param("id")
			db.Where("id=?", id).First(&td)
			c.BindJSON(&td)
			db.Save(&td)
			c.JSON(http.StatusOK, td)
		})
	}

	r.Run(":8080")
}
