package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		// 方法一：
		// 获取URL携带的 query string 参数
		name := c.Query("query")

		// 方法二：
		//name:=c.DefaultQuery("query","没查到")

		// 方法三：
		//name,ok:=c.GetQuery("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})
   
	r.Run(":8080")
}
