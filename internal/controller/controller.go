package controller

import (
	"SeKai/internal/controller/api"
	"github.com/gin-gonic/gin"
)

func InitController(router *gin.Engine) {
	pingController(router)
	homeController(router)
	backStageController(router)
	staticController(router)
	api.APIController(router)
}
