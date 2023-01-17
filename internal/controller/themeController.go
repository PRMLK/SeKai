package controller

import (
	"SeKai/internal/logger"
	"SeKai/internal/themeLoader"
	"github.com/gin-gonic/gin"
)

func themeController(router *gin.Engine) {
	for _, page := range themeLoader.BackStageTheme.Pages {
		router.GET(page.ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(page.CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
	for _, page := range themeLoader.FrontStageTheme.Pages {
		router.GET(page.ControllerURL, func(context *gin.Context) {
			_, err := context.Writer.Write(page.CompileString)
			if err != nil {
				logger.ServerLogger.Error(err)
				return
			}
		})
	}
}
