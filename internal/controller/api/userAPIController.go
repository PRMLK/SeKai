package api

import (
	"SeKai/internal/service"
	"github.com/gin-gonic/gin"
)

func userAPIController(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/login", service.LoginService)
	}
}
