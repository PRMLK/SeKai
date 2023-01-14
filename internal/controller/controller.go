package controller

import (
	"SeKai/internal/controller/api"
	"github.com/gin-gonic/gin"
)

func InitController(router *gin.Engine) {
	pingController(router)
	homeController(router)
	staticController(router)
	api.APIController(router)
}
