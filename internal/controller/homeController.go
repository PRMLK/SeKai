package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homeController(router *gin.Engine) {
	router.LoadHTMLGlob("./themes/frontstage/plain-sekai/home/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "Main website",
		})
	})
}
