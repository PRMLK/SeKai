package themeLoader

import (
	"SeKai/internal/logger"
	"github.com/pelletier/go-toml/v2"
	"os"
)

func ThemeBasicScan(dir string, themeMap map[string]themeConfig) {
	// 扫描backStage目录
	backStageRootDir, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, dir := range backStageRootDir {
		// 如果是目录
		if dir.IsDir() {
			var data []byte
			var themeConfig themeConfig
			// 读取该目录下manifest.toml
			if tempData, err := os.ReadFile("./themes/backStage/" + dir.Name() + "/manifest.toml"); err != nil {
				logger.ServerLogger.Debug()
				continue
			} else {
				data = tempData
			}
			// 尝试读取到Config
			if err := toml.Unmarshal(data, &themeConfig); err != nil {
				logger.ServerLogger.Debug()
				continue
			} else {
				if _, exist := themeMap[themeConfig.ThemeName]; exist == true {
					// 已经存在同名模板
					logger.ServerLogger.Debug()
					continue
				} else {
					themeConfig.ThemeDir = dir.Name()
					themeMap[themeConfig.ThemeName] = themeConfig
				}
			}
		}
	}
}

func SingleThemeScan(themeName string, themeMap map[string]themeConfig) {

}