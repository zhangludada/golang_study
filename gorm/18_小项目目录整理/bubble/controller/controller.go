package controller

import (
	"main/dao"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	todo.Status = true
	c.BindJSON(&todo)
	dao.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func Todolist(c *gin.Context) {
	var title []models.Todo
	dao.DB.Where("status=?", true).Find(&title)
	c.JSON(http.StatusOK, title)
}

func TodoAlist(c *gin.Context) {
	var td models.Todo
	id := c.Param("id")
	dao.DB.Where("id=?", id).First(&td)
	c.JSON(http.StatusOK, td)
}

func DeleteATodo(c *gin.Context) {
	var td models.Todo
	id := c.Param("id")
	dao.DB.Where("id=?", id).First(&td)
	dao.DB.Model(&td).Update("status", false)
	c.JSON(http.StatusOK, td)
}

func EditATodo(c *gin.Context) {
	var td models.Todo
	id := c.Param("id")
	dao.DB.Where("id=?", id).First(&td)
	c.BindJSON(&td)
	dao.DB.Save(&td)
	c.JSON(http.StatusOK, td)
}
