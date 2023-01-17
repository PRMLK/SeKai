package main

import (
	"SeKai/internal/config"
	"SeKai/internal/logger"
	"SeKai/internal/server/http"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	logger.ServerLogger.Info(config.LanguageConfig.ServerLogger.LoadConfigMessage)
	logger.ServerLogger.Info(config.LanguageConfig.ServerLogger.HTTPStartingMessage)
	http.StartHTTP()
}
