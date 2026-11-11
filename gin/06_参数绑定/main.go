package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	// tag 用 form 让 Shouldbind 找到字段
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// 自动绑定结构体字段
		var u UserInfo
		c.ShouldBind(&u)

		c.JSON(http.StatusOK, u)
	})

	r.POST("/", func(c *gin.Context) {
		// 自动绑定结构体字段
		var u UserInfo
		c.ShouldBind(&u)
		c.JSON(http.StatusOK, u)
	})

	r.Run(":8080")
}
