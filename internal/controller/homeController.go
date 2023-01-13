package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homeController(router *gin.Engine) {
	// 暂时的解决方法
	router.LoadHTMLFiles(
		"./internal/chunkLoader/chunk.tmpl",
		"./themes/frontStage/plain-sekai/pages/home/content.tmpl",
		"./themes/frontStage/plain-sekai/pages/post/content.tmpl",
		"./themes/frontStage/plain-sekai/template/footer/footer.tmpl",
		"./themes/frontStage/plain-sekai/template/header/header.tmpl",
		"./themes/frontStage/plain-sekai/template/mask/mask.tmpl",
	)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chunk", gin.H{
			"sekaiPageTitle": "Main website",
		})
	})
}
