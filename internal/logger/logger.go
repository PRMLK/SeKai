package logger

import (
	"SeKai/internal/config"
	"os"
)

var logStream *os.File

func InitLogger() {
	// 初始化日志文件对象
	if tempLogStream, err := os.OpenFile(config.ApplicationConfig.Log.Dir, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		panic("Open log file error!")
	} else {
		logStream = tempLogStream
	}
	initGinLogger()
	initServerLogger()
}
