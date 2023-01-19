package themeLoader

import (
	"SeKai"
	"SeKai/internal/config"
	"SeKai/internal/logger"
)

func inlineTemplateStringLoader() string {
	file, err := SeKai.InlineTmpl.ReadFile("internal/themeLoader/tmpl/root.tmpl")
	if err != nil {
		logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.InlineFileReadError, err)
		return ""
	}
	return string(file)
}
