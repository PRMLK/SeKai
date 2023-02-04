package api

import (
	"SeKai/internal/middleware/api"
	"SeKai/internal/service"
	"github.com/gin-gonic/gin"
)

func postAPIController(router *gin.RouterGroup) {
	post := router.Group("/post")
	{
		post.POST("/new", api.AuthMiddleware(), service.NewPost)
		post.PUT("/:id", api.AuthMiddleware(), service.EditPost)
		post.GET("/:id", service.ShowPost)
		post.DELETE("/:id", service.DelPost)
		post.GET("/list", service.GetPostList)
	}
}
