package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 1. 中间件
// 2. 登陆状态参数传递

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})

	// 接收是否登录
	flag := c.MustGet("login")
	fmt.Printf("已登陆：%v\n", flag)
}

// 中间件，统计耗时
func m1(c *gin.Context) {
	// 记时
	now := time.Now()
	// 后续的处理函数
	c.Next()
	cost := time.Since(now)
	fmt.Println("cost:", cost)

}

func m2(c *gin.Context) {
	fmt.Println("m2 in")
	// 设置登录状态
	c.Set("login", true)
	// c.Abort() 阻止后续的函数
	c.Next()
	fmt.Println("m2 out")
}

func main() {
	r := gin.Default()

	// 全局注册中间件函数
	r.Use(m1, m2)

	// 第二个参数可以接受多个函数（中间件）
	r.GET("/", m1, indexHandler)

	r.Run(":8080")
}
