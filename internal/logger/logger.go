package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
	"time"
)

var Gin = logrus.New()
var LogEntry *logrus.Entry

/*
初始化Logger

Refer： https://blog.csdn.net/zy_whynot/article/details/120240327 zy_whynot 于 2021-09-11 17:16:11 发布
*/
func init() {
	// Logger格式
	Gin.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%] %time% - %msg% \n",
	})

	// 日志存放路径
	logPath := "logs/log"

	// 最新日志的软连接路径
	linkName := "logs/latest.log"

	// 初始化日志文件对象
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// 把产生的日志内容写进日志文件中
	Gin.Out = io.MultiWriter(src, os.Stdout)

	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	// 设置日志级别
	Gin.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotatelogs.New(
		// 日志文件名格式
		logPath+"%Y%m%d.log",
		// 最多保留7天之内的日志
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 一天保存一个日志文件
		rotatelogs.WithRotationTime(24*time.Hour),
		// 为最新日志建立软连接
		rotatelogs.WithLinkName(linkName),
	)

	// 使用logWriter写日志
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// Hook格式
	Hook := lfshook.NewHook(writeMap, &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%] %time% - %msg%\n",
	})

	Gin.AddHook(Hook)
	LogEntry = logrus.NewEntry(Gin).WithField("service", "yi-shou-backstage")
}
