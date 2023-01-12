package logger

import (
	"SeKai/internal/config"
	"os"
)

var logStream *os.File

func InitLogger() {
	if _, err := os.Stat(config.ApplicationConfig.Log.Dir); err != nil {
		if err := os.Mkdir("./logs", 0755); err != nil {
			panic("Create logs dictionary error!")
		}
	}
	// 初始化日志文件对象
	if tempLogStream, err := os.OpenFile(config.ApplicationConfig.Log.Dir+"/log", os.O_RDWR|os.O_CREATE, 0755); err != nil {
		panic("Open log file error!")
	} else {
		logStream = tempLogStream
	}
	initHTTPLogger()
	initServerLogger()
}
