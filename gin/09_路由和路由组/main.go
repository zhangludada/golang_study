package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 一组路由
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			{
				c.JSON(200, gin.H{"method": "get"})
			}
		case http.MethodPost:
			{
				c.JSON(200, gin.H{"method": "post"})
			}
		case http.MethodPut:
			{
				c.JSON(200, gin.H{"method": "put"})
			}
		case http.MethodDelete:
			{
				c.JSON(200, gin.H{"method": "delete"})
			}
		}
	})

	// 路由组
	videGroup := r.Group("/video")
	videGroup.GET("/index/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"status": "ok",
			"id":     id})
	})

	r.Run(":8080")
}
