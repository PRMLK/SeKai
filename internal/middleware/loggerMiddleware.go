package middleware

import (
	"SeKai/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

/*
Logger的Gin中间件

Refer： https://blog.csdn.net/zy_whynot/article/details/120240327 zy_whynot 于 2021-09-11 17:16:11 发布
*/
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 调用该请求的剩余处理程序
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		//hostName, err := os.Hostname()
		//if err != nil {
		//    hostName = "Unknown"
		//}
		statusCode := c.Writer.Status()
		//clientIP := c.ClientIP()
		//userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := logger.GinLogger.WithFields(logrus.Fields{
			//"HostName": hostName,
			"SpendTime": spendTime,
			"path":      url,
			"method":    method,
			"status":    statusCode,
			//"Ip": clientIP,
			//"DataSize": dataSize,
			//"UserAgent": userAgent,
		})
		// 创建内部错误
		if len(c.Errors) > 0 {
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode >= 500 {
			Log.Error()
		} else if statusCode >= 400 {
			Log.Warn()
		} else {
			Log.Info()
		}
	}
}
