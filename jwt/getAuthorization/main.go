package main

import (
    "github.com/gin-gonic/gin"

    "main/handlers"
    "main/middleware"
)

func main() {

    r := gin.Default()

    // 无需认证的登录接口
    r.POST("/login", handlers.LoginHandler)

    // 需要认证的路由组
    api := r.Group("/api")

    api.Use(middleware.AuthMiddleware())

    // 需要登录才能访问
    api.GET("/profile", handlers.ProfileHandler)

    r.Run(":8080")
}