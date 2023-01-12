package controller

import "github.com/gin-gonic/gin"

func InitController(router *gin.Engine) {
	pingController(router)
}
