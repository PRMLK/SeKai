package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homeController(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chunk.tmpl", gin.H{
			"sekaiPageTitle": "Home",
			"sekaiSiteRoot":  "http://localhost:12070",
		})
	})
}
