package main

import (
	"SeKai/internal/config"
	"SeKai/internal/logger"
	"SeKai/internal/server/http"
)

func main() {
	logger.Logger.Info("SeKai starting...")
	config.InitConfig()
	logger.Logger.Info("Load config successful...")
	http.StartHTTP()
}
