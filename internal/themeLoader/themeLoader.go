package themeLoader

import (
	"SeKai/internal/config"
	"SeKai/internal/model"
)

var backStageThemeMap map[string]themeConfig
var frontStageThemeMap map[string]themeConfig

var BackStageTheme model.Theme
var FrontStageTheme model.Theme

func InitThemeLoader() {
	backStageThemeMap = make(map[string]themeConfig)
	frontStageThemeMap = make(map[string]themeConfig)
	//BackStageTheme = model.Theme{}
	//FrontStageTheme = model.Theme{}

	ThemeBasicScan("./themes/backStage", backStageThemeMap)
	ThemeBasicScan("./themes/frontStage", frontStageThemeMap)

	SingleThemeScan(
		"./themes/backStage",
		config.ApplicationConfig.SiteConfig.SiteBackStageTheme,
		backStageThemeMap,
		&BackStageTheme,
	)
	SingleThemeScan(
		"./themes/frontStage",
		config.ApplicationConfig.SiteConfig.SiteFrontStageTheme,
		frontStageThemeMap,
		&FrontStageTheme,
	)
}
