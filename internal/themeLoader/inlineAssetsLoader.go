package themeLoader

import (
	"SeKai"
	"SeKai/internal/config"
	"SeKai/internal/logger"
)

func inlineAssetsLoader(templateMap map[string]string) map[string]string {
	if tempByte, err := SeKai.InlineTmpl.ReadFile("internal/themeLoader/tmpl/chunk.tmpl"); err != nil {
		logger.ServerLogger.Fatal(config.LanguageConfig.ServerLogger.ChunkTemplateLoadedError)
	} else {
		templateMap["chunk"] = string(tempByte)
	}
	return templateMap
}
