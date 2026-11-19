package models

import "github.com/golang-jwt/jwt/v5"

// Claims 用于定义 JWT Payload 中存储的数据
type Claims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`

    // 继承 JWT 标准声明，例如过期时间 ExpiresAt
    jwt.RegisteredClaims
}