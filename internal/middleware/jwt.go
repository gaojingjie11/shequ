package middleware

import (
	"smartcommunity/pkg/response"
	"smartcommunity/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Header 中的 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.FailWithCode(c, 401, "请先登录")
			c.Abort() // 阻止后续处理
			return
		}

		// 2. 格式校验 (通常是 "Bearer <token>")
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.FailWithCode(c, 401, "Token格式错误")
			c.Abort()
			return
		}

		// 3. 解析 Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.FailWithCode(c, 401, "Token无效或已过期")
			c.Abort()
			return
		}

		// 4. 将当前用户ID存入 Context，供后续 Handler 使用
		// 【关键】后续在 Handler 里用 c.GetInt64("userID") 取出来
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next() // 放行
	}
}
