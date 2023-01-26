package controller

import (
	"SeKai/internal/controller/api"
	"github.com/gin-gonic/gin"
)

func InitController(router *gin.Engine) {
	// 主题中配置的路由
	themeController(router)
	// 主题中配置的静态文件
	staticController(router)
	// api中的路由
	api.APIController(router)
}
