package http

import (
	"SeKai/internal/controller"
	"SeKai/internal/middleware"
	"github.com/gin-gonic/gin"
)

func StartHTTP() {
	router := gin.New()
	// 加载中间件
	middleware.LoadMiddleware(router)
	// 加载控制器
	controller.InitController(router)
	router.Run(":12070")
}
