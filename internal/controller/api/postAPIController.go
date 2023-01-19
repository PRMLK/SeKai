package api

import (
	"SeKai/internal/middleware/api"
	"SeKai/internal/service"
	"github.com/gin-gonic/gin"
)

func postAPIController(router *gin.RouterGroup) {
	post := router.Group("/post")
	{
		post.POST("/new", api.AuthMiddleware(), service.NewPostService)
		post.PUT("/:id", api.AuthMiddleware(), service.EditPostService)
		post.GET("/:id", service.ShowPostService)
	}
}
