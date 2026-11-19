package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"

    "main/models"
)

var jwtSecret = []byte("replace-this-with-a-long-random-secret")

func AuthMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        // 获取 Authorization Header
        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Missing Authorization header",
            })
            return
        }

        // 分割 Bearer 和 Token
        parts := strings.SplitN(authHeader, " ", 2)

        if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Invalid Authorization format",
            })
            return
        }

        tokenString := parts[1]

        // 解析和验证 JWT
        token, err := jwt.ParseWithClaims(
            tokenString,
            &models.Claims{},
            func(token *jwt.Token) (interface{}, error) {

                // 只允许 HS256
                if token.Method != jwt.SigningMethodHS256 {
                    return nil, jwt.ErrSignatureInvalid
                }

                return jwtSecret, nil
            },
        )

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Invalid or expired token",
            })
            return
        }

        // 提取 Claims
        claims, ok := token.Claims.(*models.Claims)

        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Invalid token claims",
            })
            return
        }

        // 将用户信息存入 Gin Context
        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)

        // 放行
        c.Next()
    }
}