package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"

    "main/models"
)

var jwtSecret = []byte("replace-this-with-a-long-random-secret")

// LoginHandler 登录接口
func LoginHandler(c *gin.Context) {

    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    // 解析前端提交的 JSON
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid request",
        })
        return
    }

    // 模拟数据库验证
    if req.Username != "zl" || req.Password != "123456" {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "Wrong username or password",
        })
        return
    }

    // 创建 JWT Claims
    claims := models.Claims{
        UserID: 1001,
        Role:   "admin",
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(
                time.Now().Add(2 * time.Hour),
            ),
            IssuedAt: jwt.NewNumericDate(time.Now()),
        },
    }

    // 创建 JWT Token
    token := jwt.NewWithClaims(
        jwt.SigningMethodHS256,
        claims,
    )

    // 使用 Secret 签名
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Token generation failed",
        })
        return
    }

    // 返回 Token 给前端
    c.JSON(http.StatusOK, gin.H{
        "message": "Login success",
        "token":   tokenString,
    })
}

// ProfileHandler 获取用户信息
func ProfileHandler(c *gin.Context) {

    // 从 Middleware 中获取用户 ID
    userID, _ := c.Get("user_id")
    role, _ := c.Get("role")

    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome",
        "user_id": userID,
        "role":    role,
    })
}