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
	// 储存所有主题的manifest.toml中的信息
	backStageThemeMap = make(map[string]themeConfig)
	frontStageThemeMap = make(map[string]themeConfig)

	// 扫描所有主题的基本信息
	ThemeBasicScan("./themes/backStage", backStageThemeMap)
	ThemeBasicScan("./themes/frontStage", frontStageThemeMap)

	// 加载配置文件中指定的主题
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
