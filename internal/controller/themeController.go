package controller

import (
	"SeKai/internal/logger"
	"SeKai/internal/themeLoader"
	"github.com/gin-gonic/gin"
)

func themeController(router *gin.Engine) {
	for _, entrance := range themeLoader.BackStageTheme.Entrances {
		router.GET(entrance.ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(entrance.CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
	for _, entrance := range themeLoader.FrontStageTheme.Entrances {
		router.GET(entrance.ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(entrance.CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
}
