package main

import (
	"SeKai/internal/config"
	"SeKai/internal/logger"
	"SeKai/internal/server/http"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	logger.ServerLogger.Info("Load config and logger successful.")
	logger.ServerLogger.Info("Now starting HTTP server...")
	http.StartHTTP()
}
