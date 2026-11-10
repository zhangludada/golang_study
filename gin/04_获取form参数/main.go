package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 这里用terminal运行 go build && ./main
// 因为路径原因
func main() {
	r := gin.Default()
	// get 请求
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 获取 from 内容
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")       // 方法一
		password, _ := c.GetPostForm("password") // 方法二
		c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	r.Run(":8080")
}
