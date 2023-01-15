package api

import (
	response "SeKai/internal/controller/api/utils"
	"SeKai/internal/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, nil, "请求的auth为空")
			c.Abort()
			return
		}

		_, reqClaims, err := util.ParseToken(authHeader)
		if err != nil {
			response.Unauthorized(c, nil, "token验证失败")
			c.Abort()
			return
		}
		c.Set("userId", reqClaims.UserId)
		c.Next()

	}
}
