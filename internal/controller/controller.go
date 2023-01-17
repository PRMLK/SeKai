package controller

import (
	"SeKai/internal/controller/api"
	"github.com/gin-gonic/gin"
)

func InitController(router *gin.Engine) {
	themeController(router)
	staticController(router)
	api.APIController(router)
}
