package http

import (
	"SeKai/internal/chunkLoader"
	"SeKai/internal/config"
	"SeKai/internal/controller"
	"SeKai/internal/logger"
	"SeKai/internal/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func StartHTTP() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 加载中间件
	middleware.LoadMiddleware(router)
	// 加载themesLoader
	chunkLoader.InitLoader(router)
	// 加载控制器
	controller.InitController(router)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.ApplicationConfig.HTTP.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.ServerLogger.Panic("HTTP服务器开启失败：" + err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.ServerLogger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
