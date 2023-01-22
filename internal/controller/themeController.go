package controller

import (
	"SeKai/internal/logger"
	"SeKai/internal/themeLoader"
	"github.com/gin-gonic/gin"
)

func themeController(router *gin.Engine) {
	for i := range themeLoader.BackStageTheme.Entrances {
		router.GET(themeLoader.BackStageTheme.Entrances[i].ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(themeLoader.BackStageTheme.Entrances[i].CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
	for i := range themeLoader.FrontStageTheme.Entrances {
		router.GET(themeLoader.FrontStageTheme.Entrances[i].ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(themeLoader.FrontStageTheme.Entrances[i].CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
}
