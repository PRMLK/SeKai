package themeLoader

import (
	"github.com/gin-gonic/gin"
)

var ThemeMap map[string]themeConfig

func InitThemeLoader(router *gin.Engine) {
	ThemeMap = make(map[string]themeConfig)
	backStageThemeScan()
}
