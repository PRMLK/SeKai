package http

import (
	"SeKai/internal/config"
	"SeKai/internal/controller"
	"SeKai/internal/listener"
	"SeKai/internal/logger"
	"SeKai/internal/middleware"
	"SeKai/internal/themeLoader"
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 加载中间件
	middleware.LoadMiddleware(router)
	// 加载themesLoader
	themeLoader.InitThemeLoader()
	// 加载控制器
	controller.InitController(router)
	// 加载监听器
	listener.InitListener()

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.ApplicationConfig.HTTP.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.ServerLogger.Panic(config.LanguageConfig.ServerLogger.HTTPStartingError + ": " + err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.ServerLogger.Info(config.LanguageConfig.ServerLogger.HTTPServerShutdownMessage)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(config.LanguageConfig.ServerLogger.HTTPServerShutdownError+": ", err)
	}
	log.Println(config.LanguageConfig.ServerLogger.HTTPServerExited)
}
