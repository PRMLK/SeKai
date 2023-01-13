package http

import (
	"SeKai/internal/chunkLoader"
	"SeKai/internal/config"
	"SeKai/internal/controller"
	"SeKai/internal/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StartHTTP() {
	router := gin.Default()
	// 加载中间件
	middleware.LoadMiddleware(router)
	// 加载themesLoader
	chunkLoader.InitLoader(router)
	// 加载控制器
	controller.InitController(router)
	err := router.Run(":" + strconv.Itoa(config.ApplicationConfig.HTTP.Port))
	if err != nil {
		return
	}
}
