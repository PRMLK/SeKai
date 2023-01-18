package controller

import (
	"SeKai/internal/themeLoader"
	"github.com/gin-gonic/gin"
)

func staticController(router *gin.Engine) {
	for _, staticFile := range themeLoader.BackStageTheme.StaticFiles {
		router.StaticFile(staticFile.ControllerURL, staticFile.FileDir)
	}
	for _, staticFile := range themeLoader.FrontStageTheme.StaticFiles {
		router.StaticFile(staticFile.ControllerURL, staticFile.FileDir)
	}
}
