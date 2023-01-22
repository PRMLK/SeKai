package controller

import (
	"SeKai/internal/logger"
	"SeKai/internal/themeLoader"
	"github.com/gin-gonic/gin"
)

func themeController(router *gin.Engine) {
	for i := range themeLoader.BackStageTheme.Entrances {
		nowEntrance := &themeLoader.BackStageTheme.Entrances[i]
		router.GET(themeLoader.BackStageTheme.Entrances[i].ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(nowEntrance.CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
	for i := range themeLoader.FrontStageTheme.Entrances {
		nowEntrance := &themeLoader.FrontStageTheme.Entrances[i]
		router.GET(themeLoader.FrontStageTheme.Entrances[i].ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(nowEntrance.CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
}
