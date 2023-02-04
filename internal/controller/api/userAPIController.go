package api

import (
	"SeKai/internal/logger"
	"SeKai/internal/middleware/api"
	"SeKai/internal/service"
	"github.com/gin-gonic/gin"
)

func userAPIController(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/login", service.LoginService)
		user.POST("/register", service.RegisterService)
		user.POST("/setGoogleAuthSecret", api.AuthMiddleware(), service.SetGoogleAuthSecret)
		router.GET("/ping", api.AuthMiddleware(), func(c *gin.Context) {
			userId := c.MustGet("userId").(string)
			_, err := c.Writer.WriteString(userId)
			if err != nil {
				logger.ServerLogger.Warning(err)
				return
			}
		})
		user.GET("/profile", api.AuthMiddleware(), service.Profile)
		user.PUT("/profile", api.AuthMiddleware(), service.UpdateProfile)

	}
}
