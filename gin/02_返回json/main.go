package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 方法一
	// 用gin.H{} 设置返回值
	data := gin.H{
		"name":    "小王子",
		"message": "hello world",
		"age":     18}

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})

	// 方法二
	// 用结构体设置返回值
	type msg struct {
		Name    string
		Message string
		Age     int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"小王子",
			"hello golang",
			18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
