package controller

import (
	"SeKai/internal/logger"
	"SeKai/internal/middleware/api"
	"github.com/gin-gonic/gin"
)

func pingController(router *gin.Engine) {
	router.GET("backstage/ping", api.AuthMiddleware(), func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		_, err := c.Writer.WriteString(userId)
		if err != nil {
			logger.ServerLogger.Warning(err)
			return
		}
	})
}
