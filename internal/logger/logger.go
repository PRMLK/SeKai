package logger

import (
	"SeKai/internal/config"
	"os"
)

var logStream *os.File

func InitLogger() {
	if fileInfo, err := os.Stat(config.ApplicationConfig.Log.Dir); err != nil || !fileInfo.IsDir() {
		if err := os.Mkdir(config.ApplicationConfig.Log.Dir, 0755); err != nil {
			panic("Create logs directory error!")
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
