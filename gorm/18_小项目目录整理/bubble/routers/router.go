package routers

import (
	"main/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 创建一个路由组
	v1Group := r.Group("/v1")
	{
		// 添加一个代办事项
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有代办事项
		v1Group.GET("/todo", controller.Todolist)

		// 查看一个代办事项
		v1Group.GET("/todo/:id", controller.TodoAlist)

		// 删除一个代办事项
		// 仅修改 status = false
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

		// 编辑一个代办事项
		v1Group.PUT("/todo/:id", controller.EditATodo)
	}
	return r
}
