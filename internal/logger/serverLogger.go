package logger

import (
	"SeKai/internal/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
	"time"
)

var ServerLogger = logrus.New()

/*
初始化Logger

Refer： https://blog.csdn.net/zy_whynot/article/details/120240327 zy_whynot 于 2021-09-11 17:16:11 发布
*/
func initServerLogger() {
	// Logger格式
	ServerLogger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[ServerLogger][%lvl%] %time%   %msg%\n",
	})

	// 把产生的日志内容写进日志文件中
	ServerLogger.Out = io.MultiWriter(logStream, os.Stdout)

	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	// 设置日志级别
	ServerLogger.SetLevel(logrus.DebugLevel)

	GinLogger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotatelogs.New(
		// 日志文件名格式
		config.ApplicationConfig.Log.Dir+"%Y%m%d.log",
		// 最多保留7天之内的日志
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 一天保存一个日志文件
		rotatelogs.WithRotationTime(24*time.Hour),
		// 为最新日志建立软连接
		rotatelogs.WithLinkName(config.ApplicationConfig.Log.LastLogDir),
	)

	// 使用logWriter写日志
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// Hook格式
	Hook := lfshook.NewHook(writerMap, &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[ServerLogger][%lvl%] %time%   %msg%\n",
	})

	ServerLogger.AddHook(Hook)
	LogEntry = logrus.NewEntry(ServerLogger).WithField("service", "serverLogger")
}
