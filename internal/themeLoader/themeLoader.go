package themeLoader

import (
	"github.com/gin-gonic/gin"
)

var backStageThemeMap map[string]themeConfig
var frontStageThemeMap map[string]themeConfig

func InitThemeLoader(router *gin.Engine) {
	backStageThemeMap = make(map[string]themeConfig)
	frontStageThemeMap = make(map[string]themeConfig)
	ThemeBasicScan("./themes/backStage", backStageThemeMap)
	ThemeBasicScan("./themes/frontStage", frontStageThemeMap)

}
